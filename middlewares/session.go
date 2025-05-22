package middlewares

import (
	"encoding/json"
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/arifbugaresa/go-starter/utils/session"
	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) (output session.RedisData, err error) {
	token, err := GetJwtTokenFromHeader(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	// check token in session
	redisSession, err := session.RedisClient.Get(ctx, token).Result()
	if err != nil {
		response.GenerateErrorResponse(ctx, "token expired")
		return
	}

	err = json.Unmarshal([]byte(redisSession), &output)
	if err != nil {
		return
	}

	return output, nil
}
