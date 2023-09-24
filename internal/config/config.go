package config

import (
	"github.com/spf13/viper"
)

// Config for retrieving and storing system configuration
type Config struct {
	DBUrl          string `mapstructure:"DB_URL"`
	Port           string `mapstructure:"PORT"`
	Admin          string `mapstructure:"ADMIN"`
	AdminPassword  string `mapstructure:"ADMIN_PASSWORD"`
	AdminJwtSecret string `mapstructure:"ADMIN_JWT_SECRET"`
	UserJwtSecret  string `mapstructure:"USER_JWT_SECRET"`
}

var cfg Config

func LoadConfigs() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	cfg = config

	return
}

func GetConfig() Config {
	return cfg
}
