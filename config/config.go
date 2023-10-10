package config

import (
	// "fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var instance *Config
var once sync.Once 

func GetConfig() *Config {

	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(help)
		}
	})
	
	return instance
}
