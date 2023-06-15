package utils

import "github.com/spf13/viper"

func EnvVariabel(key string) (string, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		return "", err
	}

	return value, nil
}
