package common

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetPreviewURL(filePath string) string {
	if filePath != "" {
		return fmt.Sprintf("%s/get-file/%s",
			viper.GetString("app.base_url")+viper.GetString("app.port"),
			filePath,
		)
	}

	return filePath
}
