package ssh

import (
	"golang.org/x/crypto/ssh"
)

// createSshSession is a function to create an ssh.session.
func createSshSession(
	host string,
	port string,
	sshConfig *ssh.ClientConfig,
) (*ssh.Session, error) {
	var session *ssh.Session

	client, err := ssh.Dial("tcp", host+":"+port, sshConfig)
	if err != nil {
		return session, err
	}

	session, err = client.NewSession()
	if err != nil {
		return session, err
	}
	return session, nil
}
