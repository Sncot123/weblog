package redis

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"  //zset帖子发布的时间
	KeyPostScoreZSet       = "post:score" //zset帖子投票的分数
	keyPostVotedZSetPrefix = "post:voted" //zset记录用户及投票的类型;参数是post id

)

func getRedisKey(key string) string {
	return KeyPrefix + key
}

var communityMap = map[int64]string{
	1: "GO",
	2: "python",
	3: "java",
	4: "PHP",
}

func getCommunityKey(communityID int64) string {
	msg, _ := communityMap[communityID]
	return msg
}
