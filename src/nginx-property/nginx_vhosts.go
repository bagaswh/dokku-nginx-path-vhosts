package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
)

// Add these structs at the top of the file
type PropertyConfig struct {
	Name         string
	DefaultValue string
	UsesAppName  bool
}

var propertyConfigs = map[string]PropertyConfig{
	"access-log-format":       {"access-log-format", "", false},
	"access-log-path":         {"access-log-path", "", true},
	"client-body-timeout":     {"client-body-timeout", "60s", false},
	"client-header-timeout":   {"client-header-timeout", "60s", false},
	"client-max-body-size":    {"client-max-body-size", "1m", false},
	"disable-custom-config":   {"disable-custom-config", "false", false},
	"error-log-path":          {"error-log-path", "", true},
	"hsts":                    {"hsts", "true", false},
	"hsts-include-subdomains": {"hsts-include-subdomains", "true", false},
	"hsts-max-age":            {"hsts-max-age", "15724800", false},
	"hsts-preload":            {"hsts-preload", "false", false},
	"keepalive-timeout":       {"keepalive-timeout", "75s", false},
	"lingering-timeout":       {"lingering-timeout", "5s", false},

	"app-path":              {"app-path", "", true},
	"root-domain":           {"root-domain", "", true},
	"default-app":           {"default-app", "", true},
	"strip-path":            {"strip-path", "false", false},
	"underscore-in-headers": {"underscore-in-headers", "off", false},

	"nginx-conf-location-sigil-path":     {"nginx-conf-location-sigil-path", "location.conf.sigil", false},
	"nginx-conf-upstream-sigil-path":     {"nginx-conf-upstream-sigil-path", "upstream.conf.sigil", false},
	"nginx-conf-server-block-sigil-path": {"nginx-conf-server-block-sigil-path", "server-block.conf.sigil", false},
	"nginx-conf-http-block-sigil-path":   {"nginx-conf-http-block-sigil-path", "http-block.conf.sigil", false},

	"proxy-buffer-size":       {"proxy-buffer-size", "4k", false},
	"proxy-buffering":         {"proxy-buffering", "on", false},
	"proxy-buffers":           {"proxy-buffers", "8 4k", false},
	"proxy-busy-buffers-size": {"proxy-busy-buffers-size", "8k", false},
	"proxy-connect-timeout":   {"proxy-connect-timeout", "60s", false},
	"proxy-read-timeout":      {"proxy-read-timeout", "60s", false},
	"proxy-send-timeout":      {"proxy-send-timeout", "60s", false},
	"x-forwarded-for-value":   {"x-forwarded-for-value", "$proxy_add_x_forwarded_for", false},
	"x-forwarded-port-value":  {"x-forwarded-port-value", "$proxy_port", false},
	"x-forwarded-proto-value": {"x-forwarded-proto-value", "$scheme", false},
	"x-forwarded-ssl":         {"x-forwarded-ssl", "on", false},
}

var wildcardProperties = map[string]PropertyConfig{
	"nginx-conf-location-sigil-path-*":     {"", "", false},
	"nginx-conf-upstream-sigil-path-*":     {"", "", false},
	"nginx-conf-server-block-sigil-path-*": {"", "", false},
	"nginx-conf-http-block-sigil-path-*":   {"", "", false},
}

func lookupProperty(property string) (PropertyConfig, bool) {
	for wildcardProperty, prop := range wildcardProperties {
		if strings.HasPrefix(property, strings.Trim(wildcardProperty, "*")) {
			prop.Name = property
			return prop, true
		}
	}

	prop, exists := propertyConfigs[property]
	return prop, exists
}

// Generic property getters
func GetAppProperty(appName string, property PropertyConfig) string {
	return common.PropertyGet(getProxyName(), appName, property.Name)
}

func GetComputedProperty(appName string, property string) string {
	prop, exists := lookupProperty(property)
	if !exists {
		return ""
	}

	appValue := GetAppProperty(appName, prop)
	if appValue != "" {
		return appValue
	}

	return GetGlobalProperty(appName, prop)
}

func GetGlobalProperty(appName string, property PropertyConfig) string {
	if property.UsesAppName {
		return common.PropertyGetDefault(getProxyName(), "--global", property.Name, getDefaultValue(property, appName))
	}
	return common.PropertyGetDefault(getProxyName(), "--global", property.Name, property.DefaultValue)
}

func getDefaultValue(config PropertyConfig, appName string) string {
	switch config.Name {
	case "access-log-path":
		return fmt.Sprintf("%s/%s-access.log", getLogRoot(), appName)
	case "error-log-path":
		return fmt.Sprintf("%s/%s-error.log", getLogRoot(), appName)
	default:
		return config.DefaultValue
	}
}

func getProxyName() string {
	if v := os.Getenv("PROXY_NAME"); v != "" {
		return v
	}
	return "nginx"
}
