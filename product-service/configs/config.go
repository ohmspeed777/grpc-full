package configs

import (
	"fmt"
	"strings"

	"github.com/ohmspeed777/go-pkg/logx"
	"github.com/spf13/viper"
)

type MongoDB struct {
	Database string
	Username string
	Password string
	Host     string
	Port     int
}

func (m *MongoDB) ToURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", m.Username, m.Password, m.Host, m.Port)
}

type Config struct {
	MongoDB MongoDB
}

func NewConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		logx.GetLog().Fatalf("cannot read in viper config:%s", err)
	}

	return &Config{
		MongoDB: MongoDB{
			Database: viper.GetString("mongo.database"),
			Username: viper.GetString("mongo.username"),
			Password: viper.GetString("mongo.password"),
			Host:     viper.GetString("mongo.host"),
			Port:     viper.GetInt("mongo.port"),
		},
	}
}
