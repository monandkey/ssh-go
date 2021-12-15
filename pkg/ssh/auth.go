package ssh

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func sshPasswordAuthorization(userName string, password string) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

func sshPublicKeyAuthorization(userName string, publicKey string) (*ssh.ClientConfig, error) {
	var sshconfig *ssh.ClientConfig
	buf, err := ioutil.ReadFile(publicKey)
	if err != nil {
		return sshconfig, err
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return sshconfig, err
	}

	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	return config, nil
}
