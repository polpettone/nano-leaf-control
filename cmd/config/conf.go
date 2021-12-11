package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const nanoLeafURL = "nano_leaf_url"
const token = "token"

func GetURL(nanoLeafID int64) string {
	nanoLeafURL := viper.GetString(fmt.Sprintf("url_%d", nanoLeafID))
	token := viper.GetString(fmt.Sprintf("token_%d", nanoLeafID))
	apiPath := "api/v1"
	url := fmt.Sprintf("http://%s/%s/%s", nanoLeafURL, apiPath, token)
	return url
}
