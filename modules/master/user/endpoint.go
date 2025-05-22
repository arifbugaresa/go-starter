package user

import (
	"github.com/arifbugaresa/go-starter/middlewares"
	"github.com/arifbugaresa/go-starter/utils/common/message"
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB, redisConnection *redis.Client) {
	var (
		userRepo = NewRepository(dbConnection)
		userSrv  = NewService(userRepo, redisConnection)
	)

	publicAPI := router.Group("v1/api")
	{
		// v1/api/register
		publicAPI.POST("/register", func(c *gin.Context) {
			RegisterEndpoint(c, userSrv)
		})
		// v1/api/login
		publicAPI.POST("/login", func(c *gin.Context) {
			LoginEndpoint(c, userSrv)
		})
	}

	// user api
	protectedUserApi := router.Group("v1/api/users")
	protectedUserApi.Use(middlewares.JwtUserMiddleware())
	{
		// v1/api/users
		protectedUserApi.GET("", func(c *gin.Context) {
			GetProfileEndpoint(c, userSrv)
		})
		// v1/api/users
		protectedUserApi.PUT("", func(c *gin.Context) {
			UpdateProfileEndpoint(c, userSrv)
		})
		// v1/api/users/logout
		protectedUserApi.POST("/logout", func(c *gin.Context) {
			LogoutEndpoint(c, userSrv)
		})
	}
}

func RegisterEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody RegisterRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateRegisterRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	result, err := userSrv.RegisterUser(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, message.RegisterSuccessful, result)
}

func LoginEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody LoginRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateLoginRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	result, err := userSrv.Login(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, message.LoginSuccessful, result)
}

func LogoutEndpoint(ctx *gin.Context, userSrv Service) {
	err := userSrv.Logout(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponse(ctx, message.LogoutSuccessful)
}

func GetProfileEndpoint(ctx *gin.Context, userSrv Service) {
	record, err := userSrv.GetUserProfile(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, message.GetProfileSuccessful, record)
}

func UpdateProfileEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody UpdateProfileRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateUpdateProfileRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err = userSrv.UpdateUserProfile(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponse(ctx, message.UpdateProfileSuccessful)
}
