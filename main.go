package main

import (
	"os"

	"github.com/monandkey/ssh/pkg/ssh"
	"github.com/spf13/cobra"
)

type params struct {
	host      string
	port      string
	user      string
	password  string
	publicKey string
	command   string
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
		host:      "",
		port:      "22",
		user:      "",
		password:  "",
		publicKey: "",
		command:   "",
	}

	rootCmd.Use = "ssh"
	rootCmd.Short = "ssh command test"
	rootCmd.Version = "0.1"
	rootCmd.Flags().StringVarP(&params.host, "host", "H", params.host, "")
	rootCmd.Flags().StringVarP(&params.port, "port", "p", params.port, "")
	rootCmd.Flags().StringVarP(&params.user, "user", "u", params.user, "")
	rootCmd.Flags().StringVarP(&params.password, "password", "P", params.password, "")
	rootCmd.Flags().StringVarP(&params.publicKey, "identity-file", "i", params.publicKey, "")
	rootCmd.Flags().StringVarP(&params.command, "command", "c", params.command, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.host == "" || params.user == "" {
			return rootCmd.Help()
		}

		actour := ssh.SshStrct()
		actour.Set(
			params.host,
			params.port,
			params.user,
			params.password,
			params.publicKey,
			params.command,
		)

		config, err := actour.Authentication()
		if err != nil {
			return err
		}

		session, err := actour.Connect(config)
		if err != nil {
			return err
		}
		defer session.Close()

		if err := actour.Run(session); err != nil {
			return err
		}

		return nil
	}
}

func main() {
	Execute()
}
