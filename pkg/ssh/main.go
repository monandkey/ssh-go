package ssh

import (
	"github.com/monandkey/ssh/pkg/log"
	"golang.org/x/crypto/ssh"
)

func SshStrct(host []string) SshMethod {
	if len(host) == 1 {
		return singleNodeStruct()
	}
	return multiNodeStruct()
}

func (s *sshConfig) Set(
	host []string,
	port []string,
	user []string,
	password []string,
	publicKey []string,
	command string,
) {
	s.host = host
	s.port = port
	s.user = user
	s.password = password
	s.publicKey = publicKey
	s.command = command
}

func (s *sshConfig) Authentication() ([]*ssh.ClientConfig, error) {
	var (
		clientConfig []*ssh.ClientConfig
		user         = s.user[0]
		publicKey    = s.publicKey[0]
		password     = s.password[0]
	)

	if publicKey != "" {
		cfg, err := sshPublicKeyAuthorization(user, publicKey, password)
		if err != nil {
			return clientConfig, err
		}
		clientConfig = append(clientConfig, cfg)
	} else {
		cfg, err := sshPasswordAuthorization(user, password)
		if err != nil {
			return clientConfig, err
		}
		clientConfig = append(clientConfig, cfg)
	}
	return clientConfig, nil
}

func (s *sshConfig) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var (
		host = s.host[0]
		port = s.port[0]
		cfg  = sshConfig[0]
	)

	var sessions []*ssh.Session
	session, err := createSshSession(host, port, cfg)
	if err != nil {
		return sessions, err
	}
	sessions = append(sessions, session)
	return sessions, nil
}

func (s *sshConfig) Run(sessions []*ssh.Session) error {
	var host = s.host[0]
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if s.command == "" {
			if err := interactiveShellCalling(session); err != nil {
				return err
			}
		} else {
			logger := loggerFactory.NewLogger(host)
			nonInteractiveShellCalling(session, s.command, logger)
		}
	}
	return nil
}
