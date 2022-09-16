package models

//bluebell/models/param.go

// ParamSignUp 注册信息
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`    //必填 注册账号
	Password   string `json:"password" binding:"required"`    //必填 注册密码
	RePassword string `json:"re_password" binding:"required"` // 必填 重复密码
}

// ParamLogin 登录
type ParamLogin struct {
	UserName string `json:"username" binding:"required"` //必填 账号
	Password string `json:"password" binding:"required"` // 必填 密码
}

const ( // 注意：一般常量写在最上面
	OrderTime  = "time"
	OrderScore = "score"
)

// paramPostList2 帖子列表的请求信息
type ParamPostsList2 struct {
	Page  int64  `form:"page"`  //请求帖子列表的页码
	Size  int64  `form:"size"`  //请求帖子列表的数量
	Order string `form:"order"` //请求帖子列表的排序依据
}

// ParamPostsListCommunity 按社区请求帖子的的列表
type ParamPostsListCommunity struct {
	*ParamPostsList2       //帖子列表的请求信息
	CommunityId      int64 `json:"community_id"` //帖子的社区信息
}
