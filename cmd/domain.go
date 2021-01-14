package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ebrahimahmadi/ar-cli/pkg/http"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

type SearchResponse struct {
	Data []DomainData
}

type InfoResponse struct {
	Data []DomainData
}

type ListNSRecordResponse struct {
	Data struct {
		NSDomain []string `json:"ns_domain"`
		NSKeys   []string `json:"ns_keys"`
	}
}

type CheckNSResponse struct {
	Message string `json:"message"`
	Data    struct {
		Status bool `json:"ns_status"`
	}
}

type DomainData struct {
	UUID      string            `json:"id"`
	Name      string            `json:"name"`
	Domain    string            `json:"domain"`
	Services  map[string]string `json:"services"`
	Status    string            `json:"status"`
	NSKeys    []string          `json:"ns_keys"`
	CurrentNS []string          `json:"current_ns"`
}

var DomainName string
var DomainId string

var descriptions = map[string]string{
	"command":         "Create, Search, Delete, Get, Health check and get Ns records ",
	"search":          "Leaving the 'search' flag is empty, will return all domains. Otherwise, it will filter domains containing the search keyword.",
	"create":          "Create new domain",
	"info":            "Get information of the domain",
	"domain-name":     "The host name. like: example.com",
	"domain-id":       "The domain UUID. like: 3541b0ce-e8a6-42f0-b65a-f03a7c387486",
	"remove":          "Remove the domain",
	"list-ns-records": "Get list of domain's root NS records and expected values",
	"check":           "Check NS to find whether domain is activated",
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
		res, _ := http.Get("https://napi.arvancloud.com/cdn/4.0/domains", nil)

		responseData, _ := ioutil.ReadAll(res.Body)

		var domainInfo = new(SearchResponse)
		_ = json.Unmarshal(responseData, &domainInfo)

		table := newTable([]string{"Id", "Name", "Domain", "DNS Status", "Domain Status", "NS Key #1", "NS Key #2"})

		for _, foundDomain := range domainInfo.Data {
			record := []string{
				foundDomain.UUID,
				foundDomain.Name,
				foundDomain.Domain,
				foundDomain.Services["dns"],
				foundDomain.Status,
				foundDomain.NSKeys[0],
				foundDomain.NSKeys[1],
			}
			table.Append(record)
		}

		table.Render()

	},
}

var info = &cobra.Command{
	Use:   "info",
	Short: "get a domain info",
	Long:  descriptions["info"],
	Run: func(cmd *cobra.Command, args []string) {
		res, _ := http.Get("https://napi.arvancloud.com/cdn/4.0/domains/"+DomainName, nil)

		responseData, _ := ioutil.ReadAll(res.Body)

		var domainInfo = new(InfoResponse)
		_ = json.Unmarshal(responseData, &domainInfo)

		table := newTable([]string{"Id", "Name", "Domain", "DNS Status", "Domain Status", "NS Key #1", "NS Key #2"})

		record := []string{
			domainInfo.Data[0].UUID,
			domainInfo.Data[0].Name,
			domainInfo.Data[0].Domain,
			domainInfo.Data[0].Services["dns"],
			domainInfo.Data[0].Status,
			domainInfo.Data[0].NSKeys[0],
			domainInfo.Data[0].NSKeys[1],
		}

		table.Append(record)

		table.Render()
	},
}

var remove = &cobra.Command{
	Use:   "remove",
	Short: "remove a domain",
	Long:  descriptions["remove"],
	Run: func(cmd *cobra.Command, args []string) {
		idAsUrlQuery := map[string]string{
			"id": DomainId,
		}

		_, err := http.Delete("https://napi.arvancloud.com/cdn/4.0/domains/"+DomainName, nil, idAsUrlQuery)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Removed Successfully")
	},
}

var nsRecords = &cobra.Command{
	Use:   "ns-records",
	Short: "get list of all NS records",
	Long:  descriptions["list-ns-records"],
	Run: func(cmd *cobra.Command, args []string) {
		//TODO read from config file
		response, err := http.Get("https://napi.arvancloud.com/cdn/4.0/domains/"+DomainName+"/dns-service/ns-keys", nil)

		if err != nil {
			log.Fatal(err)
		}

		responseData, _ := ioutil.ReadAll(response.Body)

		var nsRecordList = new(ListNSRecordResponse)
		_ = json.Unmarshal(responseData, &nsRecordList)

		nsKeysTable := newTable([]string{"NS Keys"})
		nsDomainTable := newTable([]string{"NS Domain"})

		for _, nsKey := range nsRecordList.Data.NSKeys {
			record := []string{
				nsKey,
			}
			nsKeysTable.Append(record)
		}

		for _, nsDomain := range nsRecordList.Data.NSDomain {
			record := []string{
				nsDomain,
			}
			nsDomainTable.Append(record)
		}

		nsDomainTable.Render()

		nsKeysTable.Render()

	},
}

var check = &cobra.Command{
	Use:   "check",
	Short: "ensure domain is active",
	Long:  descriptions["check"],
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Put("https://napi.arvancloud.com/cdn/4.0/domains/"+DomainName+"/dns-service/check-ns", nil)

		if err != nil {
			log.Fatal(err)
		}

		responseData, _ := ioutil.ReadAll(response.Body)

		var nsStatus = new(CheckNSResponse)
		_ = json.Unmarshal(responseData, &nsStatus)

		if nsStatus.Data.Status {
			fmt.Println("NS domain is activated")
		} else {
			fmt.Println("NS domain is NOT activated")
		}

		fmt.Println(nsStatus.Message)
	},
}

func newTable(tableHeaders []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(tableHeaders)
	table.SetRowLine(true)
	table.SetRowSeparator("~")

	return table
}

func init() {
	rootCmd.AddCommand(domainCmd)

	domainCmd.AddCommand(search)
	domainCmd.AddCommand(create)
	domainCmd.AddCommand(info)
	domainCmd.AddCommand(nsRecords)
	domainCmd.AddCommand(check)
	domainCmd.AddCommand(remove)

	info.Flags().StringVarP(&DomainName, "name", "n", "", descriptions["domain-name"])

	remove.Flags().StringVarP(&DomainName, "name", "n", "", descriptions["domain-name"])
	remove.Flags().StringVarP(&DomainId, "id", "i", "", descriptions["domain-id"])

	nsRecords.Flags().StringVarP(&DomainName, "name", "n", "", descriptions["domain-name"])

	check.Flags().StringVarP(&DomainName, "name", "n", "", descriptions["domain-name"])
}
