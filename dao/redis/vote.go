package redis

import (
	"errors"
	"fmt"
	"math"
	"time"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

const (
	oneWeek      = 7 * 24 * 3600
	scorePerVote = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票已过期！")
	ErrVoteRepeated   = errors.New("重复投票")
)

func VotePost(userID, postID string, value float64) error {
	// 1、判断帖子限制
	// 去redis取帖子的发布时间 先检查帖子是否满足投票的时间
	pipeline := client.Pipeline()
	postTime := pipeline.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeek {
		return ErrVoteTimeExpire
	}
	// 2、查看该帖子的投票zset
	if pipeline.Exists(getRedisKey(keyPostVotedZSetPrefix+postID)).Val() < 1 {
		pipeline.ZAdd(getRedisKey(keyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	// 3、更新帖子分数
	//查看当前用户给帖子的投票纪录
	ov := pipeline.ZScore(getRedisKey(keyPostVotedZSetPrefix+postID), postID).Val()
	// 如果当前投票的值和之前保存的一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	fmt.Println("2")
	diff := math.Abs(ov - value) //计算两次投票的差值
	_, err := pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	fmt.Println("3")
	// 4、记录用户为该帖子投票的数据
	if value == 0 {
		_, err = pipeline.ZRem(getRedisKey(keyPostVotedZSetPrefix+postID), userID).Result()
	} else {
		_, err = pipeline.ZAdd(getRedisKey(keyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}
	fmt.Println("4")
	_, err = pipeline.Exec()
	fmt.Println("---------------------------")
	return err
}

//CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {

	//分别把帖子的id加入到社区的set、帖子的创建时间zset
	pipeline := client.Pipeline()
	// 把帖子的id和创建时间加入到时间的zset中
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: p.ID,
	})
	// 把帖子id加入到它所属的社区缓存里
	pipeline.SAdd(getCommunityKey(p.CommunityId), p.ID)
	_, err = pipeline.Exec()
	if err != nil {
		zap.L().Error("创建帖子时，把帖子id加入社区失败")
		return
	}
	return err
}
