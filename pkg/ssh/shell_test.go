package ssh

import (
	"testing"

	"github.com/monandkey/ssh/pkg/log"
	"golang.org/x/crypto/ssh"
)

func Test_nonInteractiveShellCalling(t *testing.T) {
	var (
		host     = "172.16.100.1"
		userName = "test"
		password = "test"
	)

	sshConfig := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	session, _ := createSshSession(host, "22", sshConfig)
	loggerFactory := log.NewLoggerFactory()
	logger := loggerFactory.NewLogger(host)

	type args struct {
		session *ssh.Session
		command string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				session: session,
				command: "ls",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nonInteractiveShellCalling(tt.args.session, tt.args.command, logger)
		})
	}
}
