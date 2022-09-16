package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConf struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Port        int    `mapstructure:"port"`
	StartTime   string `mapstructure:"start_time"`
	MachineID   int64  `mapstructure:"machine_id"`
	*MysqlConf  `mapstructure:"mysql"`
	*LoggerConf `mapstructure:"log"`
	*RedisConf  `mapstructure:"redis"`
}
type MysqlConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	MaxConn  int    `mapstructure:"max_size"`
	MaxIdle  int    `mapstructure:"max_idle"`
}
type LoggerConf struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
}
type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error, conf *AppConf) {
	conf = new(AppConf)
	viper.SetConfigName("config") //指定配置文件名称  (不需要带后缀)
	//viper.SetConfigType("yaml")    //指定配置文件类型 主要用来读取远程配置系统中的配置文件，从而进行特定格式解析的
	viper.AddConfigPath("./conf/") //指定查找配置文件的路径
	err = viper.ReadInConfig()     //读取配置信息
	if err != nil {
		fmt.Println("read config file failed")
		return
	}
	//把读取到的配置反序列化到conf变量中
	if err = viper.Unmarshal(conf); err != nil {
		fmt.Println("unmarshal failed error:", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file has changed")
		if err := viper.Unmarshal(conf); err != nil {
			fmt.Println("unmarshal failed error:", err)
			return
		}
	})
	return

}
