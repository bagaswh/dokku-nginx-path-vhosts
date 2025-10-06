package main

import (
	dokkuproperty "dokku-nginx-botika/src/pkg/dokku_property"
	"flag"
	"fmt"
)

var configFilePathPropertyName string = "nginx-botika-vhost-config-file"

func main() {

	var appName string
	flag.StringVar(&appName, "app", "", "the app name")

	fmt.Println("nginx-config-builder is here baby")

	configFilePath := dokkuproperty.GetComputedProperty(appName, configFilePathPropertyName)
	fmt.Println(configFilePath)

}
