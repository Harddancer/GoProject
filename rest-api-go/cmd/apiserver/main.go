package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	apiserver "github.com/Harddancer/GoProject/rest-api-go/internal/app/apiserver"
)

var (
	configPath string
)

func Init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.Newconfig()

	s := apiserver.New(config)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Start(); err != nil {
		log.Fatal(err)

	}

}
