package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Mode              string `mapstructure:"MODE" required:"true"`
	RecaptchaSecretV2 string `mapstructure:"RECAPTCHA_SECRET_V2" required:"true"`
	RecaptchaSecretV3 string `mapstructure:"RECAPTCHA_SECRET_V3" required:"true"`
}

func NewConfig() *Config {
	var cfg Config
	viper.AutomaticEnv()
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir := filepath.Dir(executablePath)
	viper.SetConfigFile(filepath.Join(executableDir, ".env"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
