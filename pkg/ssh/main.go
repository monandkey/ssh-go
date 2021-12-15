package ssh

import (
	"golang.org/x/crypto/ssh"
)

func SshStrct() SshMethod {
	return &sshConfig{}
}

func (s *sshConfig) Set(
	host string,
	port string,
	user string,
	password string,
	publicKey string,
	command string,
) {
	s.host = host
	s.port = port
	s.user = user
	s.password = password
	s.publicKey = publicKey
	s.command = command
}

func (s *sshConfig) Authentication() (*ssh.ClientConfig, error) {
	if s.password == "" {
		return sshPublicKeyAuthorization(s.user, s.publicKey)
	}
	return sshPasswordAuthorization(s.user, s.password)
}

func (s *sshConfig) Connect(sshConfig *ssh.ClientConfig) (*ssh.Session, error) {
	session, err := createSshSession(s.host, s.port, sshConfig)
	if err != nil {
		return session, err
	}
	return session, nil
}

func (s *sshConfig) Run(session *ssh.Session) error {
	if s.command == "" {
		if err := interactiveShellCalling(session); err != nil {
			return err
		}
	} else {
		if err := nonInteractiveShellCalling(session, s.command); err != nil {
			return err
		}
	}
	return nil
}
