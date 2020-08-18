package main

import (
	"flag"
	"log"

	"github.com/jialijelly/sample_blog_server/server"
)

var (
	configLocation = ""
)

func init() {
	flag.StringVar(&configLocation, "config", "config.json", "use specified config file")
	flag.Parse()

	// initialise config
	log.Printf("Loading config from %v", configLocation)
	if err := server.LoadConfigFromFile(configLocation); err != nil {
		log.Printf("Failed to load config : %v", err)
	}
	log.Printf("Using config = %+v", server.DefaultConfiguration)
}

func main() {
	server.NewServer().Run()
}
