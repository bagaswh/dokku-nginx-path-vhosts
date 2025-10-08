package main

import (
	dokkuproperty "dokku-nginx-custom/src/pkg/dokku_property"
	"flag"
	"fmt"

	"github.com/dokku/dokku/plugins/common"
)

var configFilePathPropertyName string = "config-file"

func main() {

	var appName string
	flag.StringVar(&appName, "app", "", "the app name")

	fmt.Println("nginx-config-builder is here baby")

	configFilePath := dokkuproperty.GetAppProperty(appName, configFilePathPropertyName)
	fmt.Println(configFilePath)

	fmt.Println(
		"nginx-custom=", common.PropertyGet("nginx-custom", appName, "config-file"),
		"nginx-custom-vhosts=", common.PropertyGet("nginx-custom-vhosts", appName, "config-file"),
	)

}
