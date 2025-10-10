package main

import (
	"bytes"
	"dokku-nginx-custom/src/pkg/file_config"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var configFilePathPropertyName string = "config-file"

func mustEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalln("missing required env var:", name)
	}
	return value
}

func buildUpstreamConfig(appName string, config *file_config.Config) string {

	appListeners := strings.Split(mustEnv("DOKKU_APP_LISTENERS"), " ")
	portMap := strings.Split(mustEnv("PROXY_PORT_MAP"), " ")
	upstreamPorts := strings.Split(mustEnv("PROXY_UPSTREAM_PORTS"), " ")

	templateStr := `
	{{ range $upstreamPort := split .proxyUpstreamPorts " " }} 
	upstream {{ $.app }}-{{ $upstreamPort }} {
	{{ range $listeners := split $.appListeners " " }}
	{{ $listenerList := split $listeners ":" }} 
	{{ $listenerIP := index $listenerList 0 }}
	  server {{ $listenerIP }}:{{ $upstreamPort }};{{ end }}
	}
	{{ end }}
	`

	tmplData := map[string]any{
		"app":                appName,
		"upstreamPorts":      upstreamPorts,
		"appListeners":       appListeners,
		"proxyUpstreamPorts": portMap,
	}

	tmpl, err := template.New("").Parse(templateStr)
	if err != nil {
		log.Fatalln("failed to parse template:", err)
	}

	var tmplOut bytes.Buffer
	err = tmpl.Execute(&tmplOut, tmplData)
	if err != nil {
		log.Fatalln("failed to execute template:", err)
	}

	return tmplOut.String()
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
