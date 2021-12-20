package main

import (
	"os"

	"github.com/monandkey/ssh/pkg/ssh"
	"github.com/spf13/cobra"
)

type params struct {
	host      []string
	port      []string
	user      []string
	password  []string
	publicKey []string
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
		host:      []string{},
		port:      []string{},
		user:      []string{},
		password:  []string{},
		publicKey: []string{},
		command:   "",
	}

	rootCmd.Use = "ssh"
	rootCmd.Short = "ssh command test"
	rootCmd.Version = "0.1"
	rootCmd.SilenceUsage = true
	rootCmd.Flags().StringArrayVarP(&params.host, "host", "H", params.host, "")
	rootCmd.Flags().StringArrayVarP(&params.port, "port", "p", params.port, "")
	rootCmd.Flags().StringArrayVarP(&params.user, "user", "u", params.user, "")
	rootCmd.Flags().StringArrayVarP(&params.password, "password", "P", params.password, "")
	rootCmd.Flags().StringArrayVarP(&params.publicKey, "identity-file", "i", params.publicKey, "")
	rootCmd.Flags().StringVarP(&params.command, "command", "c", params.command, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(params.host) == 0 && len(params.user) == 0 {
			return rootCmd.Help()
		}

		add := func(array []string, inStr string) []string {
			if len(array) == 0 {
				array = append(array, inStr)
			}
			return array
		}

		addDiff := func(array, comparison []string, inStr string) []string {
			if len(array) == len(comparison) {
				return array
			}

			for i := len(array); i < len(comparison); i++ {
				array = append(array, inStr)
			}
			return array
		}

		params.host = add(params.host, "")
		params.port = addDiff(params.port, params.host, "22")
		params.user = add(params.user, "")
		params.password = add(params.password, "")
		params.publicKey = addDiff(params.publicKey, params.host, "")

		actour := ssh.SshStrct(params.host)
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

		sessions, err := actour.Connect(config)
		if err != nil {
			return err
		}

		for _, session := range sessions {
			defer session.Close()
		}

		if err := actour.Run(sessions); err != nil {
			return err
		}
		return nil
	}
}

func main() {
	Execute()
}
