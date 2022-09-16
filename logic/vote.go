package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

/*
	阅读：	关于用户投票的相关算法 www.ruanyifeng.com
	投票的几种情况：

		1、之前没投过票，现在投赞成票  -->更新分数和投票纪录
		2、之前没投过票，现在投反对票
		3、之前投反对票，现在取消
		4、之前投赞成票，现在取消
		5、之前投赞成票，现在改投反对票
		6、之前投反对票，现在改投赞成票
	1赞成  -1反对  0取消
	投票的限制
		每个帖子自发表之日起一周内允许用户投票，超过一周不允许投票
			1、到期后把帖子的赞成票和反对票存储到mysql中
			2、到期之后删除对应的KeyPostVoteZSetPrefix
*/
// PostVote 为帖子投票的函数
func PostVote(userID int64, p *models.VoteData) error {
	return redis.VotePost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
