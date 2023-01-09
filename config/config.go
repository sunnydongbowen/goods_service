package config

// @program:     goods_servies
// @file:        config.go
// @author:      bowen
// @create:      2023-01-08 20:33
// @description:

import (
	"fmt"
	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf 定义全局的变量
var Conf = new(SrvConfig)

// viper.GetXxx()读取的方式
// 注意：
// Viper使用的是 `mapstructure`

// SrvConfig 服务配置
type SrvConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	// snowflake
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`

	IP   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`

	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*ConsulConfig `mapstructure:"consul"`
}

// mysql配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// redis配置
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

// 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// consul配置
type ConsulConfig struct {
	Addr string `mapstructure:"addr"`
}

func Init(fillePath string) (err error) {
	//viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigFile(fillePath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInconfig failed,err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v", err)
	}
	// 配置文件监听
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v", err)
		}
	})
	return
}
