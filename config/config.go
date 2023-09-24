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

	Logger struct {
		WithRequestID     bool           `mapstructure:"WITH_REQUEST_ID"`
		LogFilePath       map[int]string `mapstructure:"LOG_FILE_PATH"`
		EnableWriteTxtLog bool           `mapstructure:"ENABLE_WRITE_TXT_LOG"`
		CustomLogLevels   []int          `mapstructure:"CUSTOM_LOG_LEVELS"`
		GrayScale         *struct {
			Enabled       bool `mapstructure:"ENABLED"`
			Threshold     int  `mapstructure:"THRESHOLD"`
			Percentage    int  `mapstructure:"PERCENTAGE"`
			TotalRequests int  `mapstructure:"TOTAL_REQUESTS"`
		} `mapstructure:"GRAY_SCALE"`
	} `mapstructure:"LOGGER"`
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
