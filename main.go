package main

import (
	"log"
	"net/http"
	"strconv"
)

var config *Config

func init() {
	var err error
	config, err = loadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(config.Path)))

	log.Printf("Serving %s on port %d\n", config.Path, config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
}
