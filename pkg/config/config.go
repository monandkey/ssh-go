package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// getFileName is a function to get the configuration name.
func getFileName() string {
	const configFileName string = ".ssh/ssh_config.yaml"
	homeDir := setHomedir()
	separate := setSeparate()
	return homeDir + separate + configFileName
}

// configLoad is a function for loading the configuration.
func configLoad(fileName string) []configParams {
	configParams := []configParams{}
	b, _ := os.ReadFile(fileName)
	yaml.Unmarshal(b, &configParams)
	return configParams
}
