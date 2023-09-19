package config

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/utils"
	"github.com/spf13/viper"
)

type Config struct {
	DbURL string `mapstructure:"DB_URL"`
	PORT  string `mapstructure:"PORT"`
}

func LoadConfigs() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if utils.HasError(context.Background(), err) {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
