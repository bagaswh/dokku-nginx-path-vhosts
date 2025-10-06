package main

import (
	"dokku-nginx-botika/src/pkg/file_config"
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Please provide a config file path using -config flag")
	}

	// Get query from positional argument
	args := flag.Args()
	var query string
	if len(args) > 0 {
		query = args[0]
	}

	// Read config file
	_, rawConfig, err := file_config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// If query is provided, access that specific part of config
	if query != "" {
		result, err := file_config.QueryConfig(rawConfig, query)
		if err != nil {
			log.Fatalf("Error querying config: %v", err)
		}

		// Output result as YAML
		output, err := yaml.Marshal(result)
		if err != nil {
			log.Fatalf("Error marshaling query result: %v", err)
		}
		fmt.Println(string(output))
		return
	}

	log.Fatalln("Please provide the query as positional argument")
}
