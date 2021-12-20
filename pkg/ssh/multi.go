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

// Connect is a function for creating multiple sessions.
func (m *multiNode) Connect(sshConfig *ssh.ClientConfig) ([]*ssh.Session, error) {
	var sessions []*ssh.Session
	for _, host := range m.multiHost {
		session, err := createSshSession(host, m.port, sshConfig)
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
			logger := loggerFactory.NewLogger(m.multiHost[cnt])
			nonInteractiveShellCalling(session, m.command, logger)
		} else {
			return errors.New("please specify the command")
		}
		cnt++
	}
	return nil
}
