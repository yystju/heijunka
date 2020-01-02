package main

import (
	"flag"
	"heijunka"
	"io/ioutil"
	"log"

	toml "github.com/BurntSushi/toml"
)

var (
	configFile string
	config     heijunka.Config
)

func init() {
	log.Println("[INIT]")

	flag.StringVar(&configFile, "f", "config.toml", "The config file")
	flag.Parse()

	//TODO: Check if the config file existed...

	configContent, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.Decode(string(configContent), &config)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("[MAIN]")

	log.Printf("config : %v\n", config)
}
