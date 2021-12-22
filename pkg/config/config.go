package config

import (
	"io/ioutil"
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

// configWrite is a function for writing the configuration.
func configWrite(fileName string, data interface{}) error {
	if fileExist(fileName) {
		if err := fileOpen(fileName); err != nil {
			return err
		}
	} else {
		if err := fileCreate(fileName); err != nil {
			return nil
		}
	}

	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, buf, 0664); err != nil {
		return err
	}
	return nil
}
