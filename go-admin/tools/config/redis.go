package config

import "github.com/spf13/viper"

type Redis struct {
	RedisConn string
	RedisDbNum    int
	RedisPassword string
	RedisMaxIdle int
	RedisMaxActive int
}

func InitRedis(cfg *viper.Viper) *Redis {
	return &Redis{
		RedisConn: cfg.GetString("RedisConn"),
		RedisDbNum:    cfg.GetInt("RedisDbNum"),
		RedisPassword: cfg.GetString("RedisPassword"),
		RedisMaxIdle: cfg.GetInt("RedisMaxIdle"),
		RedisMaxActive: cfg.GetInt("RedisMaxActive"),
	}
}

var RedisConfig = new(Redis)
