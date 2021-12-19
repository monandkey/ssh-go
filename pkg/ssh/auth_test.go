package ssh

import (
	"testing"

	"golang.org/x/crypto/ssh"
)

func Test_sshPasswordAuthorization(t *testing.T) {
	var (
		username = "root"
		password = "r00t"
	)

	type args struct {
		userName string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.ClientConfig
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				userName: username,
				password: password,
			},
			want: &ssh.ClientConfig{
				User: username,
				Auth: []ssh.AuthMethod{
					ssh.Password(password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sshPasswordAuthorization(tt.args.userName, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("sshPasswordAuthorization() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
