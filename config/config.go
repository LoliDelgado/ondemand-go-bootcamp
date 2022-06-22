package config

import "github.com/spf13/viper"

// Config stores all configuration of the application
type Config struct {
	Port           string `mapstructure:"PORT" validate:"required"`
	SourceFilePath string `mapstructure:"SOURCE_FILE_PATH" validate:"required"`
	SourceFileName string `mapstructure:"SOURCE_FILE_NAME" validate:"required"`
}

// LoadConfig reads configuration from file or environment variables
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
