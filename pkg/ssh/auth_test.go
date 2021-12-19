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

func Test_sshPublicKeyAuthorization(t *testing.T) {
	var (
		username  = "root"
		publicKey = "/root/.ssh/id_rsa"
	)

	type args struct {
		userName   string
		publicKey  string
		passphrase string
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
				userName:   username,
				publicKey:  publicKey,
				passphrase: "",
			},
			wantErr: false,
		},
		{
			name: "Normal passphrase Test",
			args: args{
				userName:   username,
				publicKey:  "/root/.ssh/id_test",
				passphrase: "test",
			},
			wantErr: false,
		},
		{
			name: "failure Test",
			args: args{
				userName:   username,
				publicKey:  "/root/.ssh/id_abc",
				passphrase: "",
			},
			wantErr: true,
		},
		{
			name: "Failure passphrase Test",
			args: args{
				userName:   username,
				publicKey:  "/root/.ssh/id_test",
				passphrase: "testt",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sshPublicKeyAuthorization(tt.args.userName, tt.args.publicKey, tt.args.passphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("sshPublicKeyAuthorization() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
