package main

import (
	dokkuproperty "dokku-nginx-custom/src/pkg/dokku_property"
	"flag"
	"fmt"
	"os"

	"github.com/dokku/dokku/plugins/common"
)

var configFilePathPropertyName string = "config-file"

func main() {

	fmt.Println("args=", os.Args)

	var appName string
	flag.StringVar(&appName, "app", "", "the app name")
	if appName == "" {
		// log.Fatalln("--app flag is required")
	}

	configFilePath := dokkuproperty.GetAppProperty(appName, configFilePathPropertyName)
	fmt.Println(configFilePath)

	fmt.Println(
		"appName=", appName,
		"nginx=", common.PropertyGet("nginx", "laravel-app", "config-file"),
		"nginx-custom=", common.PropertyGet("nginx-custom", "laravel-app", "config-file"),
		"nginx-custom-vhosts=", common.PropertyGet("nginx-custom-vhosts", "laravel-app", "config-file"),
	)

}
