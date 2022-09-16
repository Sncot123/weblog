package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routes"
	"web_app/settings"

	"go.uber.org/zap"
)

// @title bluebell这是我写的第一个项目
// @version 1.0
// @description 这是我李某人些的第一个web后端，我去，忒激动了
// @termsOfService http://swagger.io/terms/

// @contact.name 红木资本
// @contact.url http://www.swagger.io/support
// @contact.email sncot123@aliyun.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	// 1、加载配置
	err, conf := settings.Init()
	if err != nil {
		fmt.Printf("init settings failed  error:", err)
		return
	}
	// 2、初始化日志
	if err := logger.Init(conf.LoggerConf, conf.Mode); err != nil {
		fmt.Printf("init logger failed  error:", err)
		return
	}
	defer zap.L().Sync() //把内存中的日志写入磁盘
	zap.L().Debug("logger init successfully...")
	// 3、初始化mysql连接
	if err := mysql.Init(conf.MysqlConf); err != nil {
		fmt.Printf("init logger failed  error:", err)
		return
	}
	defer mysql.Close()
	// 4、初始化redis连接
	if err := redis.Init(conf.RedisConf); err != nil {
		fmt.Printf("init logger failed  error:", err)
		return
	}
	defer redis.Close()
	if err = snowflake.Init(conf.StartTime, conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed error:%s\n", err)
		return
	}
	// 5、注册路由
	router := routes.SetUp(conf.Mode)
	// 6、启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}
	go func() {
		//开启一个goroutine开启服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen:%s\n", zap.Error(err))
		}
	}()
	//等待信号来优雅的关闭服务器，为关闭服务器设置一个5秒的延迟
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	//kill会默认发送syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在这 当收到上述两种信号时才会往下执行
	log.Println("shutdown server...")
	//擦黄建一个5秒的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel() //5秒内优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown  ", zap.Error(err))
	}
	zap.L().Info("server exiting...")
}
