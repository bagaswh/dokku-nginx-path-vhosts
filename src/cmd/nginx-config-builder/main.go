package main

import (
	dokkuproperty "dokku-nginx-custom/src/pkg/dokku_property"
	"fmt"
	"log"
	"os"
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

	configFilePath := dokkuproperty.GetAppProperty(appName, configFilePathPropertyName)
	fmt.Println("configFilePath=", configFilePath)

}
