package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1、生成post id
	p.ID = int64(snowflake.GenID())

	// 2、保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	// 3、分别把帖子的id加入到社区的set、帖子的创建时间zset、帖子得分的zset
	err = redis.CreatePost(p)
	if err != nil {
		return
	}
	return
}

func GetPostById(id int64) (data *models.ApiPost, err error) {

	// 查询并组合我们想用的数据
	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostById(id) failed", zap.Int64("pid", id), zap.Error(err))
		return
	}

	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Error(err))
		return
	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetail(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(post.CommunityId)", zap.Error(err))
		return
	}
	//拼装信息
	data = &models.ApiPost{
		AuthorName:      user.UserName,
		CommunityName:   community.Name,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
