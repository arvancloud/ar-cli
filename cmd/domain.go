package cmd

import (
	"encoding/json"
	"github.com/ebrahimahmadi/ar-cli/pkg/http"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type SearchResponse struct {
	Data []struct{
		Name string `json:"name"`
		Domain string `json:"domain"`
		Status string `json:"status"`
	}
}

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
		// todo: READ FROM CONFIG FILE
		// todo: Add search parameter
		res, _:= http.Get("https://napi.arvancloud.com/cdn/4.0/domains", nil)

		responseData, _ := ioutil.ReadAll(res.Body)

		var searchResult = new(SearchResponse)
		_ = json.Unmarshal(responseData, &searchResult)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Domain", "Status"})

		for _, foundDomain := range searchResult.Data {
			record := []string{
				foundDomain.Name,
				foundDomain.Domain,
				foundDomain.Status,
			}
			table.Append(record)
		}

		table.SetRowLine(true)
		table.SetRowSeparator("~")
		table.Render()

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
