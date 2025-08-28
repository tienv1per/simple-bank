package util

import "github.com/spf13/viper"

// Config stores all configuration of application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	if path != "" {
		viper.AddConfigPath(path)
	}

	// Set default values
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("SERVER_ADDRESS", "127.0.0.1:8080")

	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
