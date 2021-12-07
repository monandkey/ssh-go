package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type params struct {
	host     string
	port     string
	user     string
	password string
}

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func init() {
	params := params{
		host:     "",
		port:     "22",
		user:     "",
		password: "",
	}

	rootCmd.Use = "ssh"
	rootCmd.Short = "ssh command test"
	rootCmd.Version = "0.1"
	rootCmd.Flags().StringVarP(&params.host, "host", "H", params.host, "")
	rootCmd.Flags().StringVarP(&params.port, "port", "p", params.port, "")
	rootCmd.Flags().StringVarP(&params.user, "user", "u", params.user, "")
	rootCmd.Flags().StringVarP(&params.password, "password", "P", params.password, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.host == "" || params.user == "" {
			return rootCmd.Help()
		}

		// Create sshClientConfig
		sshConfig := &ssh.ClientConfig{
			User: params.user,
			Auth: []ssh.AuthMethod{
				ssh.Password(params.password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		// SSH connect.
		client, err := ssh.Dial("tcp", params.host+":"+params.port, sshConfig)
		if err != nil {
			return err
		}

		// Create Session
		session, err := client.NewSession()
		if err != nil {
			return err
		}
		defer session.Close()

		// キー入力を接続先が認識できる形式に変換する(ここがキモ)
		fd := int(os.Stdin.Fd())
		state, err := term.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer term.Restore(fd, state)

		// ターミナルサイズの取得
		w, h, err := term.GetSize(fd)
		if err != nil {
			return err
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          1,
			ssh.TTY_OP_ISPEED: 14400,
			ssh.TTY_OP_OSPEED: 14400,
		}

		if err := session.RequestPty("xterm", h, w, modes); err != nil {
			return err
		}

		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Stdin = os.Stdin

		if err := session.Shell(); err != nil {
			return err
		}

		if err := session.Wait(); err != nil {
			return err
		}

		return nil
	}
}

func main() {
	Execute()
}
