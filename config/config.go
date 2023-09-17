package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var cfg Config
var doOnce sync.Once

type Config struct {
	Application struct {
		Port int `mapstructure:"PORT"`
	} `mapstructure:"APPLICATION"`

	MovieApi struct {
		Token string `mapstructure:"TOKEN"`
	} `mapstructure:"MOVIE_API"`
}

func Get() Config {
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(fmt.Sprintf("cannot read .env file: %v", err))
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("cannot unmarshaling config")
		}
	})

	return cfg
}
