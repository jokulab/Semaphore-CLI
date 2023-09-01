package cmd

import (
	"fmt"	
	"bytes"
	"encoding/json"
	"net/http"

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
	fmt.Println("Try login...")

	url := "http://localhost:3000/api/auth/login"
	data := map[string]string{
		"auth":     "admin",
		"password": "admin",
	}
	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

}