package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

func GetPostList(page, size int64) (data []*models.ApiPost, err error) {
	var (
		posts     []*models.Post
		user      *models.User
		community *models.CommunityDetail
	)
	posts, err = mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList()", zap.Error(err))
		return
	}

	for _, post := range posts {
		// 查询并组合我们想用的数据
		post, err = mysql.GetPostById(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetPostById(id) failed", zap.Error(err))
			return
		}
		//根据作者id查询作者信息
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Error(err))
			return
		}
		//根据社区id查询社区详细信息
		community, err = mysql.GetCommunityDetail(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail(post.CommunityId)", zap.Error(err))
			return
		}
		//拼装信息
		p := &models.ApiPost{
			AuthorName:      user.UserName,
			CommunityName:   community.Name,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, p)
	}
	return
}

//GetPostsList2 升级版 以指定顺序查询帖子列表
func GetPostsList2(postList2 models.ParamPostsList2) (data []*models.ApiPost, err error) {
	var (
		ids      []string
		voteData []int64
	)
	if ids, err = redis.GetPostsIDSOrderBy(postList2); err != nil {
		zap.L().Error("redis.GetPostsIDSOrderBy(p) 以指定序列从redis中获取ids失败")
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostsIDSOrderBy(p) return 0 data")
		return
	}
	// 提前查询好每篇帖子的投票数
	voteData, err = redis.GetPostsVoteData(ids)
	if err != nil {
		return
	}
	// 根据获取的ids去数据库查询帖子的详细信息
	var posts []*models.Post
	if posts, err = mysql.GetPostList2(ids); err != nil {

	}
	var (
		user      *models.User
		community *models.CommunityDetail
	)
	// 将帖子的作者及分区信息查询出来 填充到帖子中
	for idx, post := range posts {
		// 查询并组合我们想用的数据
		post, err = mysql.GetPostById(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetPostById(id) failed", zap.Error(err))
			return
		}
		//根据作者id查询作者信息
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Error(err))
			return
		}
		//根据社区id查询社区详细信息
		community, err = mysql.GetCommunityDetail(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail(post.CommunityId)", zap.Error(err))
			return
		}
		//拼装信息
		p := &models.ApiPost{
			AuthorName:      user.UserName,
			CommunityName:   community.Name,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, p)
	}
	return

}
