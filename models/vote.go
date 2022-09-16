package models

//bluebell/models/vote.go
//关于投票的

//VoteData 投票的数据
type VoteData struct {
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=0 1 -1"` //投票类型
}
