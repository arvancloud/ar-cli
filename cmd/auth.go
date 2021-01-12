package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	authDesc = ` Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Log in to Arvan server",
	Long:  authDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
