package main

import (
	dokkuproperty "dokku-nginx-custom/src/pkg/dokku_property"
	"dokku-nginx-custom/src/pkg/file_config"
	"fmt"
	"log"
	"os"
	"path"
)

var configFilePathPropertyName string = "config-file"

func buildUpstreamConfig(appName string, config *file_config.Config) string {
	resultCfg := ""

	// first, build default upstream retrieved from env vars
	appListeners := os.Getenv("DOKKU_APP_LISTENERS")
	portMap := os.Getenv("PROXY_PORT_MAP")
	upstreamPorts := os.Getenv("PROXY_UPSTREAM_PORTS")

	fmt.Println("appListeners:", appListeners)
	fmt.Println("portMap:", portMap)
	fmt.Println("upstreamPorts:", upstreamPorts)

	// 	templateStr := `
	// {{ if $.DOKKU_APP_WEB_LISTENERS }}
	// {{ range $upstream_port := $.PROXY_UPSTREAM_PORTS | split " " }}
	// upstream {{ $.APP }}-{{ $upstream_port }} {
	// {{ range $listeners := $.DOKKU_APP_WEB_LISTENERS | split " " }}
	// {{ $listener_list := $listeners | split ":" }}
	// {{ $listener_ip := index $listener_list 0 }}
	//   server {{ $listener_ip }}:{{ $upstream_port }};{{ end }}
	// }
	// {{ end }}{{ end }}
	// `

	return resultCfg
}

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

	cfg, rawCfg, err := file_config.ReadConfig(configFilePath)
	if err != nil {
		log.Fatalln("error parsing config file:", err)
	}
	_ = cfg
	_ = rawCfg

	upstreamCfg := buildUpstreamConfig(appName, cfg)
	fmt.Println(upstreamCfg)

}
