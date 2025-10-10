package main

import (
	"dokku-nginx-custom/src/pkg/file_config"
	"flag"
	"fmt"
	"log"
	"os"
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

	var appName string
	var configFilePath string
	flag.StringVar(&appName, "app-name", "", "app name")
	flag.StringVar(&configFilePath, "config-file-path", "", "path to config file")

	flag.Parse()

	required := []string{"app-name", "config-file-path"}

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			log.Fatalf("missing required -%s argument/flag", req)
		}
	}

	cfg, rawCfg, err := file_config.ReadConfig(configFilePath)
	if err != nil {
		log.Fatalln("error parsing config file:", err)
	}
	_ = cfg
	_ = rawCfg

	upstreamCfg := buildUpstreamConfig(appName, cfg)
	fmt.Println(upstreamCfg)

}
