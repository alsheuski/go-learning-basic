package api

import "purple.learn/3-struct/config"

func GetAPIKey() string {
	config := config.NewConfig()
	return config.Key
}
