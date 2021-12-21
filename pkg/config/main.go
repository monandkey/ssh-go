/*
This package is for configuration management.
*/
package config

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
	return nil
}

// GetConfigName is a function to get the configuration name.
func GetConfigName() string {
	return getFileName()
}
