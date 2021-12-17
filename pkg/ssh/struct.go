package ssh

import "golang.org/x/crypto/ssh"

type SshMethod interface {
	Set(string, []string, string, string, string, string, string)
	Authentication() (*ssh.ClientConfig, error)
	Connect(*ssh.ClientConfig) ([]*ssh.Session, error)
	Run([]*ssh.Session) error
}

type sshConfig struct {
	singleHost string
	multiHost  []string
	port       string
	user       string
	password   string
	publicKey  string
	command    string
}

type singleNode struct {
	sshConfig
}

type multiNode struct {
	sshConfig
}
