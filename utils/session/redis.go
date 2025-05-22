package session

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strconv"
)

type RedisData struct {
	Id       int64   `json:"user_id"`
	Photo    *string `json:"photo"`
	FullName string  `json:"full_name"`
	UserName string  `json:"username"`
	Role     string  `json:"role"`
	RoleId   int64   `json:"role_id"`
	Email    string  `json:"email"`
}

var (
	RedisClient *redis.Client
)

func Initiator() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("connection.redis.port" + ":" + strconv.Itoa(viper.GetInt("connection.redis.port"))),
		Password: viper.GetString("connection.redis.password"),
		DB:       viper.GetInt("connection.redis.db"),
	})

	return RedisClient
}
