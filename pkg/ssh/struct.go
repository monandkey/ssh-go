package ssh

import "golang.org/x/crypto/ssh"

//	SshMethod is a method that summarizes the steps to make an ssh connection.
type SshMethod interface {
	Set([]string, []string, []string, []string, []string, string)
	Authentication() ([]*ssh.ClientConfig, error)
	Connect([]*ssh.ClientConfig) ([]*ssh.Session, error)
	Run([]*ssh.Session) error
}

// sshConfig is a structure that manages the parameters required for an ssh connection.
type sshConfig struct {
	host      []string
	port      []string
	user      []string
	password  []string
	publicKey []string
	command   string
}

// singleNode is a structure for defining methods for a single node
type singleNode struct {
	sshConfig
}

// singleNode is a structure for defining methods for a multi node
type multiNode struct {
	sshConfig
}
