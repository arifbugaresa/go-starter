package upload

import "mime/multipart"

type (
	UploadFileRequest struct {
		File       *multipart.FileHeader
		ModuleName string `json:"module_name"`
	}

	UploadFileResponse struct {
		FilePath   string `json:"file_path"`
		PreviewURL string `json:"preview_url"`
	}
)

type (
	GetFileResponse struct {
		FilePath   string `json:"file_path"`
		PreviewURL string `json:"preview_url"`
	}
)
