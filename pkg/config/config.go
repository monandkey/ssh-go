package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

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

func configExampleLoad() []configParams {
	cfg := []configParams{
		{
			Name:     "server with jump",
			User:     "appuser",
			Host:     "192.168.8.35",
			Port:     "22",
			Password: "123456",
			Keypath:  "",
			Jump: []jump{
				{
					User: "appuser",
					Host: "192.168.8.36",
					Port: "2222",
				},
			},
		},
	}
	return cfg
}
