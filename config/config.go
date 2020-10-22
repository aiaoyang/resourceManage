package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GVC AliyunConfig = AliyunConfig{}

type AliyunConfig struct {
	Accounts []SingleAccount
	Regions  []string
}

type SingleAccount struct {
	Name      string
	SecretID  string
	SecretKEY string
}

func init() {
	LoadAliyunConfig()
}

func LoadAliyunConfig() {

	v := viper.New()

	v.AddConfigPath("./")

	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = v.Unmarshal(&GVC)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(GVC)

	v.WatchConfig()

	run := func(in fsnotify.Event) {
		v.Unmarshal(&GVC)
	}

	v.OnConfigChange(run)

}
