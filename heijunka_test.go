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

func TestHeijunka0(t *testing.T) {
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

	categories := make(map[string][]string)

	categories["g1"] = []string {"a"}
	categories["g2"] = []string {"b"}

	h := NewHeijunka(&config, orders, categories)

	h.Process()

	log.Printf("orders : %v, categories : %v, result : %v", orders, categories, h.Items)
}

func TestHeijunka1(t *testing.T) {
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

	orders["a"] = 10
	orders["b"] = 1

	categories := make(map[string][]string)

	categories["g1"] = []string {"a"}
	categories["g2"] = []string {"b"}

	h := NewHeijunka(&config, orders, categories)

	h.Process()

	log.Printf("orders : %v, categories : %v, result : %v", orders, categories, h.Items)
}

func TestHeijunka2(t *testing.T) {
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
	orders["c"] = 3

	categories := make(map[string][]string)

	categories["g1"] = []string {"a"}
	categories["g2"] = []string {"b"}
	categories["g3"] = []string {"c"}

	h := NewHeijunka(&config, orders, categories)

	h.Process()

	log.Printf("orders : %v, categories : %v, result : %v", orders, categories, h.Items)
}

func TestHeijunka3(t *testing.T) {
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
	orders["c"] = 3

	categories := make(map[string][]string)

	categories["g1"] = []string {"a"}
	categories["g2"] = []string {"b", "c"}

	h := NewHeijunka(&config, orders, categories)

	h.Process()

	log.Printf("orders : %v, categories : %v, result : %v", orders, categories, h.Items)
}