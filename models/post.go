package models

import "time"

//bluebell/models/post.go

// Post  创建帖子
type Post struct {
	Title       string    `json:"title" db:"title" binding:"required"`               //必填 帖子的标题
	Content     string    `json:"content" db:"content" binding:"required"`           //必填 帖子的内容
	ID          int64     `json:"id" db:"post_id"`                                   //自动生成 帖子的id
	AuthorId    int64     `json:"author_id" db:"author_id"`                          //帖子的作者id
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"` //必填 帖子所属社区id
	Status      int32     `json:"status" db:"status"`                                //选填 帖子的状态
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      //帖子的创建时间
}

// ApiPost 返回帖子的详细信息
type ApiPost struct {
	AuthorName       string `json:"author_name"`    //作者姓名
	CommunityName    string `json:"community_name"` //社区名称
	VoteNum          int64  `json:"votedate"`       //帖子的投票数量
	*Post                   //嵌入帖子信息
	*CommunityDetail        //嵌入社区信息
}
