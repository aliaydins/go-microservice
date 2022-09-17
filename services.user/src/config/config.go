package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	DbDriver       string `mapstructure:"DB_DRIVER"`
	DbUser         string `mapstructure:"DB_USER"`
	DbPassword     string `mapstructure:"DB_PASSWORD"`
	DbPort         string `mapstructure:"DB_PORT"`
	DbHost         string `mapstructure:"DB_HOST"`
	DbName         string `mapstructure:"DB_NAME"`
	AppPort        string `mapstructure:"APP_PORT"`
	SecretKey      string `mapstructure:"SECRET_KEY"`
	RabbitUser     string `mapstructure:"RABBIT_USER"`
	RabbitPassword string `mapstructure:"RABBIT_PASSWORD"`
	RabbitHost     string `mapstructure:"RABBIT_HOST"`
	RabbitPort     string `mapstructure:"RABBIT_PORT"`
	UserExchange   string `mapstructure:"USER_EXCHANGE"`
}

func (c Config) GetDBURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DbHost, c.DbPort, c.DbUser, c.DbName, c.DbPassword)
}

func (c Config) GetRabbitMQURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.RabbitUser, c.RabbitPassword, c.RabbitHost, c.RabbitPort)
}
func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config
	return AppConfig

}
