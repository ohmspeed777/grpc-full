package main

import (
	"app/cmd/protocol"
	"strings"

	"github.com/ohmspeed777/go-pkg/logx"
	"github.com/spf13/viper"
)

func main() {
	logx.Init("info", true)
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		logx.GetLog().Fatalf("cannot read in viper config:%s", err)
	}

	protocol.Init()
	// go protocol.NewGRPC()
	protocol.NewAPI()
}
