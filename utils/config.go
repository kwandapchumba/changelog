package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString     string        `mapstructure:"connection_string"`
	ConnectionPort       string        `mapstructure:"connection_port"`
	Secret               string        `mapstructure:"secret"`
	Hex                  string        `mapstructure:"hex"`
	AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
