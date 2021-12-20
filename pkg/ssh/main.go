package ssh

import (
	"github.com/monandkey/ssh/pkg/log"
	"golang.org/x/crypto/ssh"
)

func SshStrct(singleHost string) SshMethod {
	if singleHost == "" {
		return multiNodeStruct()
	}
	return singleNodeStruct()
}

func (s *sshConfig) Set(
	singleHost string,
	multiHost []string,
	port string,
	user string,
	password string,
	publicKey string,
	command string,
) {
	s.singleHost = singleHost
	s.multiHost = multiHost
	s.port = port
	s.user = user
	s.password = password
	s.publicKey = publicKey
	s.command = command
}

// Authentication is a function used to create a configuration for authentication.
func (s *sshConfig) Authentication() (*ssh.ClientConfig, error) {
	if s.password == "" {
		return sshPublicKeyAuthorization(s.user, s.publicKey, s.password)
	}
	return sshPasswordAuthorization(s.user, s.password)
}

// Connect is a function for creating multiple sessions.
func (s *sshConfig) Connect(sshConfig *ssh.ClientConfig) ([]*ssh.Session, error) {
	var sessions []*ssh.Session
	session, err := createSshSession(s.singleHost, s.port, sshConfig)
	if err != nil {
		return sessions, err
	}
	sessions = append(sessions, session)
	return sessions, nil
}

// Run is a function that sends a command to multiple devices.
func (s *sshConfig) Run(sessions []*ssh.Session) error {
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if s.command == "" {
			if err := interactiveShellCalling(session); err != nil {
				return err
			}
		} else {
			logger := loggerFactory.NewLogger(s.singleHost)
			nonInteractiveShellCalling(session, s.command, logger)
		}
	}
	return nil
}
