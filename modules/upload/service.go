package upload

import (
	"github.com/arifbugaresa/go-starter/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Service interface {
	UploadFile(ctx *gin.Context, dataBody UploadFileRequest) (response UploadFileResponse, err error)
}

type UploadService struct {
	repo *UploadRepository
}

func NewService(repo *UploadRepository) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) UploadFile(ctx *gin.Context, dataBody UploadFileRequest) (response UploadFileResponse, err error) {
	directory := filepath.Join(viper.GetString("storage.upload.file"), dataBody.ModuleName)

	// create directory
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return
	}

	// save file
	destinationFile := filepath.Join(directory, dataBody.File.Filename)
	if err = ctx.SaveUploadedFile(dataBody.File, destinationFile); err != nil {
		return
	}

	// generate file path
	response.FilePath = filepath.Join(dataBody.ModuleName, dataBody.File.Filename)
	response.PreviewURL = common.GetPreviewURL(response.FilePath)

	return
}
