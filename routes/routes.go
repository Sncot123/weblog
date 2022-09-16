package routes

import (
	"time"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-contrib/pprof"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	_ "web_app/docs"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) (router *gin.Engine) {
	//gin.SetMode(mode) //设置模式（开发模式或发布模式）
	router = gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	router.POST("/signup", controllers.SignUpHandler)
	router.POST("/login", controllers.LoginHandler)
	router.GET("/swag/*any", gs.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", middleware.RateLimitMiddleWare(time.Second, 6), func(context *gin.Context) {
		context.String(200, "pong")
	})

	// 服务型性能调试工具
	pprof.Register(router)
	v1 := router.Group("/api/v1")
	v1.Use(middleware.JWTAuthMiddleware()) //应用中间件
	{

		v1.GET("/community", controllers.CommunityHandler)

		v1.GET("community/:id", controllers.CommunityDetailHandler)
		// 新加帖子
		v1.POST("/post", controllers.CreatePostHandler)
		// 获取某个帖子的详细信息
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		// 获取帖子列表
		v1.GET("/posts", controllers.GetPostListHandler)
		//投票
		v1.POST("/vote", controllers.PostVoteHandler)
		//获取帖子列表（升级版） 根据帖子创建时间或者帖子分数的顺序来获取帖子的排序
		v1.GET("/posts2", controllers.GetPostsListHandler2)
		// 按社区分类 根据帖子的创建时间或者帖子分数的顺序来获取帖子的排序
		v1.GET("/postByCommunity", controllers.GetCommunityPostsListHandler)

	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": 404})
	})
	return router
}
