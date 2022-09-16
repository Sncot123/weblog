package models

import "time"

//bluebell/models/community.go

// Community 社区信息
type Community struct {
	ID   int64  `json:"id" db:"community_id"`     //社区id
	Name string `json:"name" db:"community_name"` //社区名称
}

// CommunityDetail 社区详细信息
type CommunityDetail struct {
	ID          int64     `json:"id" db:"community_id"`                     //社区id
	Name        string    `json:"name" db:"community_name"`                 //社区名称
	Instruction string    `json:"introduction,omitempty" db:"introduction"` //社区介绍
	CreateTime  time.Time `json:"create_time" db:"create_time"`             //创建时间
}
