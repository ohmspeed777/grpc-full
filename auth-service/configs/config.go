package configs

import (
	"fmt"

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
	APP     APP
	JWT     JWT
	GRPC    GRPC
}

type APP struct {
	Name     string
	APIPort  int
	GRPCPort int
}

type GRPC struct {
	Order string
}

type JWT struct {
	PRIV string
	PUB  string
}

func NewConfig() *Config {
	return &Config{
		MongoDB: MongoDB{
			Database: viper.GetString("mongo.database"),
			Username: viper.GetString("mongo.username"),
			Password: viper.GetString("mongo.password"),
			Host:     viper.GetString("mongo.host"),
			Port:     viper.GetInt("mongo.port"),
		},
		APP: APP{
			APIPort:  viper.GetInt("app.api"),
			Name:     viper.GetString("app.name"),
			GRPCPort: viper.GetInt("app.grpc"),
		},
		JWT: JWT{
			PRIV: viper.GetString("jwt.key"),
			PUB:  viper.GetString("jwt.pub"),
		},
		GRPC: GRPC{
			Order: viper.GetString("grpc.order"),
		},
	}
}
