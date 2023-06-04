package models

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Mongo         = InitMongo()
	RDB           = InitRedis()
	OptionsConfig = InitConfig()
)

type Options struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Mongo struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	}

	Redis struct {
		Addr        string `yaml:"addr"`
		Port        string `yaml:"port"`
		Password    string `yaml:"password"`
		DB          int    `yaml:"DB"`
		PoolSize    int    `yaml:"poolSize"`
		MinIdleConn int    `yaml:"minIdleConn"`
	}
}

// InitConfig 读取配置文件
func InitConfig() *Options {
	optionsConfig := new(Options)
	//指定要读取的配置文件

	//文件名前缀
	viper.SetConfigName("app")
	//指定文件格式
	viper.SetConfigType("yml")
	//路径
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	//判断读取配置文件是否有误
	if err := viper.ReadInConfig(); err != nil {
		//未找到配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误
			fmt.Println("配置文件未找到: ", err)
		} else {
			fmt.Println("读取配置文件数据失败: ", err)
		}
	} else {
		//绑定数据到结构体
		if err := viper.Unmarshal(&optionsConfig); err != nil {
			fmt.Println("绑定数据失败: ", err)
		} else {
			fmt.Println("读取配置文件成功")
			return optionsConfig
		}
	}
	return optionsConfig
}

// InitMongo 初始化数据库
func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: OptionsConfig.Mongo.Username,
		Password: OptionsConfig.Mongo.Password,
	}).ApplyURI("mongodb://"+OptionsConfig.Mongo.Host+":"+OptionsConfig.Mongo.Port))
	if err != nil {
		log.Println("Connection MongoDb Error:", err)
		return nil
	}
	return client.Database(OptionsConfig.Mongo.Dbname)
}

func InitRedis() *redis.Client {
	Red := redis.NewClient(&redis.Options{
		Addr:         OptionsConfig.Redis.Addr + ":" + OptionsConfig.Redis.Port,
		Password:     OptionsConfig.Redis.Password,
		DB:           OptionsConfig.Redis.DB,
		PoolSize:     OptionsConfig.Redis.PoolSize,
		MinIdleConns: OptionsConfig.Redis.MinIdleConn,
	})

	//判断连接是否成功
	if _, err := Red.Ping(context.Background()).Result(); err != nil {
		fmt.Println("连接Redis失败", err)
	} else {

		fmt.Println("连接Redis成功")
		return Red
	}
	return Red

}
