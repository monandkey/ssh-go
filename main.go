package main

import (
	"os"

	"github.com/monandkey/ssh/pkg/ssh"
	"github.com/spf13/cobra"
)

type params struct {
	singleHost string
	port       string
	user       string
	password   string
	publicKey  string
	command    string
	multiHost  []string
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
		singleHost: "",
		port:       "22",
		user:       "",
		password:   "",
		publicKey:  "",
		command:    "",
		multiHost:  []string{},
	}

	rootCmd.Use = "ssh"
	rootCmd.Short = "ssh command test"
	rootCmd.Version = "0.1"
	rootCmd.SilenceUsage = true
	rootCmd.Flags().StringVarP(&params.singleHost, "single-host", "s", params.singleHost, "")
	rootCmd.Flags().StringVarP(&params.port, "port", "p", params.port, "")
	rootCmd.Flags().StringVarP(&params.user, "user", "u", params.user, "")
	rootCmd.Flags().StringVarP(&params.password, "password", "P", params.password, "")
	rootCmd.Flags().StringVarP(&params.publicKey, "identity-file", "i", params.publicKey, "")
	rootCmd.Flags().StringVarP(&params.command, "command", "c", params.command, "")
	rootCmd.Flags().StringArrayVarP(&params.multiHost, "multi-host", "s", params.multiHost, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.singleHost == "" || params.user == "" {
			return rootCmd.Help()
		}

		actour := ssh.SshStrct()
		actour.Set(
			params.singleHost,
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
