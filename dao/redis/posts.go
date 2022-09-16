package redis

import (
	"strconv"
	"time"
	"web_app/models"

	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

func GetPostsIDSOrderBy(p models.ParamPostsList2) (ids []string, err error) {
	//从redis中获取ids
	var key string
	if p.Order == models.OrderTime {
		key = getRedisKey(KeyPostTimeZSet)
	}
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

//GetPostsVoteData 获取帖子的投票数据
func GetPostsVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(keyPostVotedZSetPrefix + id)
	//	// 查找key中分数是1的数量----》即投赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 使用pipline 发送一次请求执行多条命令，减少RTT
	pipline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(keyPostVotedZSetPrefix + id)
		pipline.ZCount(key, "1", "1")
	}
	cmders, err := pipline.Exec()
	if err != nil {
		zap.L().Error("get vote data failed")
		return
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

//GetCommunityPostsIDSOrderBy 根据社区的分类返回已制定顺序排序的帖子
func GetCommunityPostsIDSOrderBy(p models.ParamPostsListCommunity) (ids []string, err error) {
	// 使用zinterstore把分区的帖子set与按分数排序的帖子zset生成一个新的zset
	// 按照新的zset的顺序 取数据

	// 利用缓存的key减少zinterstore的次数
	key := p.Order + strconv.Itoa(int(p.CommunityId))
	// 社区的key
	cKey := getRedisKey(getCommunityKey(p.CommunityId) + strconv.Itoa(int(p.CommunityId)))
	pipeline := client.Pipeline()
	if pipeline.Exists(key).Val() < 1 {
		// 不存在 不需要计算
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey)
		pipeline.Expire(key, 60*time.Second) // 设置过期时间
		_, err = pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	//按分数从大到小的查询
	return client.ZRevRange(key, start, end).Result()
}
