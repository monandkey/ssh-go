package ssh

import (
	"errors"

	"github.com/monandkey/ssh/pkg/log"
	"golang.org/x/crypto/ssh"
)

// multiNodeStruct is a function that returns the SshMethod
func multiNodeStruct() SshMethod {
	return &multiNode{}
}

// Authentication is a function used to create a configuration for authentication.
func (m *multiNode) Authentication() ([]*ssh.ClientConfig, error) {
	var clientConfig []*ssh.ClientConfig

	for i := 0; i < len(m.host); i++ {
		if m.publicKey[i] != "" {
			cfg, err := sshPublicKeyAuthorization(m.user[i], m.publicKey[i], m.password[i])
			if err != nil {
				return clientConfig, err
			}
			clientConfig = append(clientConfig, cfg)
		} else {
			cfg, err := sshPasswordAuthorization(m.user[i], m.password[i])
			if err != nil {
				return clientConfig, err
			}
			clientConfig = append(clientConfig, cfg)
		}
	}
	return clientConfig, nil
}

// Connect is a function for creating multiple sessions.
func (m *multiNode) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var sessions []*ssh.Session
	for i := 0; i < len(m.host); i++ {
		session, err := createSshSession(m.host[i], m.port[i], sshConfig[i])
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// Run is a function that sends a command to multiple devices.
func (m *multiNode) Run(sessions []*ssh.Session) error {
	var cnt = 0
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if m.command != "" {
			logger := loggerFactory.NewLogger(m.host[cnt])
			nonInteractiveShellCalling(session, m.command, logger)
		} else {
			return errors.New("please specify the command")
		}
		cnt++
	}
	return nil
}
