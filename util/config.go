package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config struct for the app.env file
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
	LocalAuth            bool          `mapstructure:"LOCAL_AUTH"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	LogFilePath          string        `mapstructure:"LOG_FILE_PATH"`
	LogFileName          string        `mapstructure:"LOG_FILE_NAME"`
	LogMaxSize           string        `mapstructure:"LOG_MAX_SIZE"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// Load app.env file, this file is set as environment variables in container
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
