package cmd

import (
	"github.com/ebrahimahmadi/ar-cli/pkg/http"
	"github.com/spf13/cobra"
)

var descriptions = map[string]string{
	"command": "Create, Search, Delete, Get, Health check and get Ns records ",
	"search":  "Leaving the 'search' flag is empty, will return all domains. Otherwise, it will filter domains containing the search keyword.",
	"create":  "Create new domain",
	"info":    "Get information of the domain",
	"remove":  "Remove the domain",
	"list":    "Get list of domain's root NS records and expected values",
	"check":   "Check NS to find whether domain is activated",
}

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Interact with domains",
	Long:  descriptions["command"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

var create = &cobra.Command{
	Use:   "create",
	Short: "create a domain",
	Long:  descriptions["create"],
	Run: func(cmd *cobra.Command, args []string) {
		// todo: implement validation
		payload := map[string]string{
			"domain": args[0],
		}

		// todo: READ FROM CONFIG FILE
		http.Post("https://napi.arvancloud.com/cdn/4.0/domains/dns-service", payload)
	},
}

var search = &cobra.Command{
	Use:   "search",
	Short: "search domains",
	Long:  descriptions["search"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

var info = &cobra.Command{
	Use:   "info",
	Short: "get a domain info",
	Long:  descriptions["info"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

var remove = &cobra.Command{
	Use:   "remove",
	Short: "remove a domain",
	Long:  descriptions["remove"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

var list = &cobra.Command{
	Use:   "list",
	Short: "get list of all NS records",
	Long:  descriptions["list"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement list
	},
}

var check = &cobra.Command{
	Use:   "check",
	Short: "ensure domain is active",
	Long:  descriptions["check"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)

	domainCmd.AddCommand(search)
	domainCmd.AddCommand(create)
	domainCmd.AddCommand(info)
	domainCmd.AddCommand(list)
	domainCmd.AddCommand(check)
	domainCmd.AddCommand(remove)
}
