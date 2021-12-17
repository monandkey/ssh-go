package ssh

import "golang.org/x/crypto/ssh"

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
