package util

import "github.com/spf13/viper"

type Config struct {
	DBSource   string `mapstructure:"CONN_STRING"`
	ServerName string `mapstructure:"SERVER_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
