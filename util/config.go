package util

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	DBDriverPostgres string `mapstructure:"DB_DRIVER_POSTGRES"`
	DBSourcePostgres string `mapstructure:"DB_SOURCE_POSTGRES"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig loads config from file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // path to look for the config file in
	viper.SetConfigName("app") // name of config file (without extension)
	viper.SetConfigType("env") // type of config file (without extension)
	viper.AutomaticEnv()       // read in environment variables that match

	err = viper.ReadInConfig() // find and read the config file
	if err != nil {            // if error
		return // return error
	}

	err = viper.Unmarshal(&config) // unmarshal config file into config struct
	if err != nil {                // if error
		return // return error
	}

	return // return config
}
