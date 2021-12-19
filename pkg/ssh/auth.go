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

func sshPublicKeyAuthorization(userName string, publicKey string, passphrase string) (*ssh.ClientConfig, error) {
	var sshconfig *ssh.ClientConfig
	buf, err := ioutil.ReadFile(publicKey)
	if err != nil {
		return sshconfig, err
	}

	// This function will not work properly if there is a passphrase.
	// If you have a passphrase, please use "ParsePrivateKeyWithPassphrase".
	var key ssh.Signer
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		key, err = sshPublicKeyAuthorizationWithPassphrase(buf, []byte(passphrase))
		if err != nil {
			return sshconfig, err
		}
	}

	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	return config, nil
}

func sshPublicKeyAuthorizationWithPassphrase(buf []byte, passphrase []byte) (ssh.Signer, error) {
	var signer ssh.Signer
	signer, err := ssh.ParsePrivateKeyWithPassphrase(buf, passphrase)
	if err != nil {
		return signer, err
	}
	return signer, nil
}
