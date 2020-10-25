package glob

import (
	"github.com/spf13/viper"
	"log"
)

type redisConf struct {
	Db       int
	Password string
	Host     string
	Addr     string
}

type www struct {
	Addr  string
	Host  string
	Limit int
}

var Config = struct {
	Www   www
	Redis redisConf
}{}

func InitConfig() {
	viper.SetConfigFile("app.conf")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal(err)
	}
}
