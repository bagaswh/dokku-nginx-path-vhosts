package file_config

import (
	"fmt"
	"html/template"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestResolveConfigReferences(t *testing.T) {
	fileContent, err := os.ReadFile("testdata/example.yaml")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}

	var rawConfig interface{}
	if err := yaml.Unmarshal(fileContent, &rawConfig); err != nil {
		t.Fatalf("Failed to parse YAML into config struct: %v", err)
	}

	_, _, resolveErr := ResolveConfigReferences(&config, rawConfig, map[string]map[string]any{
		"upstreams": {
			"default": "an_upstream_name",
		},
	}, template.FuncMap{
		"hello": func() string {
			return "world"
		},
	})
	if resolveErr != nil {
		t.Fatalf("Failed to resolve config references: %v", resolveErr)
	}
	fmt.Println(rawConfig)
}
