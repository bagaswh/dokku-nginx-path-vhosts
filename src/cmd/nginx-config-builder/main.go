package main

import (
	dokkuproperty "dokku-nginx-custom/src/pkg/dokku_property"
	"fmt"
	"log"
	"os"
	"path"
)

var configFilePathPropertyName string = "config-file"

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("app name (positional arg 2nd) is required")
	}

	appName := os.Args[1]

	if appName == "" {
		log.Fatalln("app name (positional arg 2nd) should be non-empty")
	}

	dataDirectory := os.Getenv("DATA_DIRECTORY")
	if dataDirectory == "" {
		log.Fatalln("DATA_DIRECTORY environment variable is required")
	}

	configFilePath := path.Join(
		dataDirectory, fmt.Sprintf("app-%s", appName), dokkuproperty.GetAppProperty(appName, configFilePathPropertyName),
	)
	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("error reading config file:", err)
	}

	fmt.Println(string(fileContent))
}
