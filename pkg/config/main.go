/*
This package is for configuration management.
*/
package config

import "fmt"

// SelectConfigUser is a function that returns an interface.
func SelectConfigUser() ConfigAction {
	return &baseParams{}
}

// SetParams is a function to set the required parameters.
func (b *baseParams) SetParams(
	name string,
	user string,
	host string,
	password string,
	keypath string,
) {
	b.name = name
	b.user = user
	b.host = host
	b.password = password
	b.keypath = keypath
}

// Load is a function for loading the configuration.
func (b *baseParams) Load() []configParams {
	fileName := getFileName()

	if !(fileExist(fileName)) {
		return configExampleLoad()
	}
	return configLoad(fileName)
}

// Write is a function for writing the configuration.
func (b *baseParams) Write() error {
	var overWriteFlag string
	fileName := getFileName()

	if fileExist(fileName) {
		fmt.Print("overwrite ? y or n: ")
		fmt.Scan(&overWriteFlag)

		if overWriteFlag != "y" {
			fmt.Println("configurations were not changed")
			return nil
		}
	}

	if err := configWrite(fileName, b.params); err != nil {
		return err
	}

	if overWriteFlag == "y" {
		fmt.Println("Overwrite!!")
	} else {
		fmt.Println("Create config file")
	}

	return nil
}

// GetConfigName is a function to get the configuration name.
func GetConfigName() string {
	return getFileName()
}
