package main

import (
	"github.com/arifbugaresa/go-starter/config"
	"github.com/arifbugaresa/go-starter/databases/connection"
	"github.com/arifbugaresa/go-starter/databases/migration"
	"github.com/arifbugaresa/go-starter/modules/health_check"
	"github.com/arifbugaresa/go-starter/modules/master/user"
	"github.com/arifbugaresa/go-starter/modules/upload"
	redisPackage "github.com/arifbugaresa/go-starter/utils/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func main() {
	config.Initiator()

	dbConnection, err := connection.Initiator()
	if err != nil {
		return
	}
	defer dbConnection.Close()

	migration.Initiator(dbConnection)
	redisConnection := redisPackage.Initiator()

	InitiateRouter(dbConnection, redisConnection)
}

func InitiateRouter(dbConnection *sqlx.DB, redisConnection *redis.Client) {
	router := gin.Default()

	corsOption := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12,
	}

	router.Use(cors.New(corsOption))

	health_check.Initiator(router)

	// module user
	user.Initiator(router, dbConnection, redisConnection)
	upload.Initiator(router, dbConnection)

	router.Run(viper.GetString("app.port"))
}
