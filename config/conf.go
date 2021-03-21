package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	DebugMode      = false
	ReadmeFilename = "README.md"
)

type flagValue struct {
	Data    string
	Message string
}

func init() {
	viper.SetEnvPrefix("webls")
	defaultValue := map[string]flagValue{
		"listen":   {":9496", "listen address"},
		"path":     {".", "path to list"},
		"sitename": {"Webls", "name display in web panel"},
		"author":   {"Webls", "copyright author display in web panel"},
		"since":    {"", "since year display in web panel"},
	}
	for k, v := range defaultValue {
		viper.SetDefault(k, v.Data)
		pflag.String(k, v.Data, v.Message)
	}
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	viper.AutomaticEnv()
}
