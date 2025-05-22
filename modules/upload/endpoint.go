package upload

import (
	"github.com/arifbugaresa/go-starter/middlewares"
	"github.com/arifbugaresa/go-starter/utils/common"
	"github.com/arifbugaresa/go-starter/utils/common/message"
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB) {
	var (
		uploadRepo = NewRepository(dbConnection)
		uploadSrv  = NewService(uploadRepo)
	)

	router.Static("/get-file", "./"+viper.GetString("storage.upload.file"))

	apiProtected := router.Group("/api/uploads")
	apiProtected.Use(middlewares.JwtMiddleware())
	{
		apiProtected.POST("", func(c *gin.Context) {
			UploadFileEndpoint(c, uploadSrv)
		})
		apiProtected.GET("", func(c *gin.Context) {
			GetFileEndpoint(c, uploadSrv)
		})
	}
}

func UploadFileEndpoint(ctx *gin.Context, uploadSrv Service) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	moduleName := ctx.PostForm("module")
	if moduleName == "" {
		response.GenerateErrorResponse(ctx, message.ModuleNameRequired)
		return
	}

	record, err := uploadSrv.UploadFile(ctx, UploadFileRequest{
		ModuleName: moduleName,
		File:       file,
	})
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, message.UploadFileSuccessful, record)
}

func GetFileEndpoint(ctx *gin.Context, uploadSrv Service) {
	filePath := ctx.Query("file_path")
	if filePath == "" {
		response.GenerateErrorResponse(ctx, message.FilePathRequired)
		return
	}

	data := GetFileResponse{
		FilePath:   filePath,
		PreviewURL: common.GetPreviewURL(filePath),
	}

	response.GenerateSuccessResponseWithData(ctx, message.GetFileSuccessful, data)
}
