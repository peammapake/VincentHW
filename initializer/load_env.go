package initializer

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	DB     DBConfig     `mapstructure:"database"`
}

type ServerConfig struct {
	ServerPort string `mapstructure:"port"`
}

type DBConfig struct {
	DBHost      string `mapstructure:"db_host"`
	DBPort      string `mapstructure:"db_port"`
	DBUser      string `mapstructure:"db_username"`
	DBPassword  string `mapstructure:"db_password"`
	DBName      string `mapstructure:"db_name"`
	DBSSLConfig string `mapstructure:"db_ssl"`
	DBTimeZone  string `mapstructure:"db_timezone"`
}

func ReadConfig() (config *Config, err error) {
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config")
	// viper.SetConfigType("yaml")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s", err))
	// }

	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
		return
	}

	return
}
