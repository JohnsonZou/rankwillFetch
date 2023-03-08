package common

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisDB *redis.Client

func InitRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redisconfig.address"),
		Password: viper.GetString("redisconfig.password"),
		DB:       0,
	})
	_, err := redisDB.Ping().Result()
	return err
}
func GetRedisDB() *redis.Client {
	return redisDB
}
