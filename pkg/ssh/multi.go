package ssh

import (
	"errors"

	"golang.org/x/crypto/ssh"
)

func multiNodeStruct() SshMethod {
	return &multiNode{}
}

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

func (m *multiNode) Run(sessions []*ssh.Session) error {
	for _, session := range sessions {
		if m.command != "" {
			if err := nonInteractiveShellCalling(session, m.command); err != nil {
				return err
			}
		} else {
			return errors.New("please specify the command")
		}
	}
	return nil
}