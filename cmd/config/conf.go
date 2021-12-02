package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const nanoLeafURL = "nano_leaf_url"
const token = "token"

func GetURL() string {
	nanoLeafURL := viper.GetString("nano_leaf_url")
	token := viper.GetString("token")
	apiPath := "api/v1"
	url := fmt.Sprintf("http://%s/%s/%s", nanoLeafURL, apiPath, token)
	return url
}
