/*
This package is for configuration management.
*/
package config

import "fmt"

// SelectConfigUser is a function that returns an interface.
func SelectConfigUser() ConfigAction {
	return &baseParams{}
}

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

func (b *baseParams) Load() {
	fileName := "~/.ssh/ssh_config.yaml"
	if true {
		fmt.Println(fileName)
	}
}

func (b *baseParams) Write() error {
	return nil
}
