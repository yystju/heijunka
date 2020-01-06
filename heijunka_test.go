package heijunka

import (
	"io/ioutil"
	"log"
	"testing"

	toml "github.com/BurntSushi/toml"
)

var (
	config Config
)

func TestHeijunka(t *testing.T) {
	configContent, err := ioutil.ReadFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.Decode(string(configContent), &config)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config : %v\n", config)

	orders := make(map[string]int)

	orders["a"] = 3
	orders["b"] = 3

	h := NewHeijunka(orders)

	log.Println(h)

	h.Process()

	log.Println(h.Items)
}
