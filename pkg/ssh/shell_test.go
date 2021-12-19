package ssh

import (
	"testing"

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
			if err := nonInteractiveShellCalling(tt.args.session, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("nonInteractiveShellCalling() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
