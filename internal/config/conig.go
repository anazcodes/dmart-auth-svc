package config

import "github.com/spf13/viper"

type Config struct {
	DbURL string `mapstructure:"DB_URL"`
	PORT  string `mapstructure:"PORT"`
}

func LoadConfigs() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigType("env")
	viper.SetConfigName("dev")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.Unmarshal(&config)
	return
}
