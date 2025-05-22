package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/arifbugaresa/go-starter/utils/common/message"
	"github.com/arifbugaresa/go-starter/utils/constant/enum"
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/arifbugaresa/go-starter/utils/session"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strings"
)

type Claims struct {
	FullName string  `json:"full_name"`
	UserName string  `json:"username"`
	Email    string  `json:"email"`
	RoleId   int64   `json:"role_id"`
	Role     string  `json:"role"`
	Photo    *string `json:"photo"`
	jwt.StandardClaims
}

func (c Claims) GenerateJwtToken() (token string, err error) {
	claims := &Claims{
		Role:           c.Role,
		Email:          c.Email,
		FullName:       c.FullName,
		UserName:       c.UserName,
		RoleId:         c.RoleId,
		Photo:          c.Photo,
		StandardClaims: jwt.StandardClaims{},
	}

	generatedTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = generatedTokenJwt.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}

	return
}

func JwtRoleMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := GetJwtTokenFromHeader(ctx)
		if err != nil {
			response.GenerateErrorResponse(ctx, err.Error())
			return
		}

		// check token in session
		redisSession, err := session.RedisClient.Get(ctx, token).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				response.GenerateErrorResponse(ctx, message.TokenExpired)
				return
			}

			response.GenerateErrorResponse(ctx, err.Error())
			return
		}

		if redisSession == "" {
			response.GenerateErrorResponse(ctx, message.InvalidToken)
			return
		}

		var authUser session.RedisData
		err = json.Unmarshal([]byte(redisSession), &authUser)
		if err != nil {
			return
		}

		if role != "" {
			if authUser.Role != role {
				response.GenerateErrorResponse(ctx, message.AccessDenied)
				return
			}
		}

		ctx.Next()
	}
}

func JwtUserMiddleware() gin.HandlerFunc {
	return JwtRoleMiddleware(enum.User)
}

func JwtAdminMiddleware() gin.HandlerFunc {
	return JwtRoleMiddleware(enum.Admin)
}

func JwtMiddleware() gin.HandlerFunc {
	return JwtRoleMiddleware("")
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return tokenString, message.AuthHeaderRequired
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, message.HeaderFormatInvalid
	}

	return parts[1], nil
}
