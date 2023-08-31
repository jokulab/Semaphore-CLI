package cmd

import (
	"fmt"
	//"net/http"
	//"os"

	"github.com/spf13/cobra"
)

type authArgs struct {
	name     string
	password string
}

var apiAuthArgs authArgs

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "Specifies the user to log in")
	loginCmd.Flags().StringP("password", "p", "", "Specifies the password to log in")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into API and store Cookie",
	Run: login,
}


func login(cmd *cobra.Command, args []string)  {
	fmt.Println("Run login")
}