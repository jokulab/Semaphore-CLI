package cmd

import (
	"fmt"
	//"net/http"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "semaphore-cli",
	Short: "Semaphore CLI is a tool to use the Semaphore API",
	Long: `Semaphore CLI is a tool to use the Ansible Semaphore API.
Source code is available at https://github.com/jokulab/semaphore-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func Execute() {
	//rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Configuration file path")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
