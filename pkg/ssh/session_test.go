package ssh

import (
	"testing"

	"golang.org/x/crypto/ssh"
)

func Test_createSshSession(t *testing.T) {
	var (
		host     = "172.16.100.1"
		userName = "test"
		password = "test"
	)

	type args struct {
		host      string
		port      string
		sshConfig *ssh.ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.Session
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				host: host,
				port: "22",
				sshConfig: &ssh.ClientConfig{
					User: userName,
					Auth: []ssh.AuthMethod{
						ssh.Password(password),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				},
			},
			wantErr: false,
		},
		{
			name: "Failure Test",
			args: args{
				host: host,
				port: "22",
				sshConfig: &ssh.ClientConfig{
					User: userName,
					Auth: []ssh.AuthMethod{
						ssh.Password("testtest"),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createSshSession(tt.args.host, tt.args.port, tt.args.sshConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSshSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
