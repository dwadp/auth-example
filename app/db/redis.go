package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			viper.GetString("db.redis.host"),
			viper.GetInt("db.redis.port"),
		),
		Password: viper.GetString("db.redis.password"),
		DB:       viper.GetInt("db.redis.index"),
	})

	return rdb
}
