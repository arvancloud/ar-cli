package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ebrahimahmadi/ar-cli/pkg/api"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

type SearchResponse struct {
	Data []DomainData
}

type InfoResponse struct {
	Data DomainData
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

var searchKeyWord string

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Interact with domains",
	Long:  helpDescriptions["domain-command"],
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement logic
	},
}

var create = &cobra.Command{
	Use:   "create",
	Short: "create a domain",
	Long:  helpDescriptions["domain-create"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			BodyPayload: map[string]string{"domain": DomainName},
			URL:         Config.GetUrl() + "/domains/dns-service",
			Method:      "POST",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		if res.StatusCode == http.StatusCreated {
			fmt.Println(DomainName + " Created Successfully")
		}
	},
}

var search = &cobra.Command{
	Use:   "search",
	Short: "search domains",
	Long:  helpDescriptions["domain-search"],
	Run: func(cmd *cobra.Command, args []string) {
		request := api.RequestBag{
			URL:        Config.GetUrl() + "/domains",
			Method:     "GET",
			URLQueries: map[string]string{"search": searchKeyWord},
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

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
	Long:  helpDescriptions["domain-info"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName,
			Method: "GET",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		responseData, _ := ioutil.ReadAll(res.Body)

		var domainInfo = new(InfoResponse)
		_ = json.Unmarshal(responseData, &domainInfo)

		table := newTable([]string{"Id", "Name", "Domain", "DNS Status", "Domain Status", "NS Key #1", "NS Key #2"})

		record := []string{
			domainInfo.Data.UUID,
			domainInfo.Data.Name,
			domainInfo.Data.Domain,
			domainInfo.Data.Services["dns"],
			domainInfo.Data.Status,
			domainInfo.Data.NSKeys[0],
			domainInfo.Data.NSKeys[1],
		}

		table.Append(record)

		table.Render()
	},
}

var remove = &cobra.Command{
	Use:   "remove",
	Short: "remove a domain",
	Long:  helpDescriptions["domain-remove"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			BodyPayload: map[string]string{"id": DomainId},
			URL:         Config.GetUrl() + "/domains/" + DomainName,
			Method:      "DELETE",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		fmt.Println("Removed Successfully")
	},
}

var nsRecords = &cobra.Command{
	Use:   "ns-records",
	Short: "get list of all NS records",
	Long:  helpDescriptions["domain-list-ns-records"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-service/ns-keys",
			Method: "GET",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		responseData, _ := ioutil.ReadAll(res.Body)

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
	Long:  helpDescriptions["domain-check"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-service/check-ns",
			Method: "PUT",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		responseData, _ := ioutil.ReadAll(res.Body)

		var nsStatus = new(CheckNSResponse)
		_ = json.Unmarshal(responseData, &nsStatus)

		fmt.Println(nsStatus.Message)
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)

	domainCmd.AddCommand(search)
	domainCmd.AddCommand(create)
	domainCmd.AddCommand(info)
	domainCmd.AddCommand(nsRecords)
	domainCmd.AddCommand(check)
	domainCmd.AddCommand(remove)

	search.Flags().StringVarP(&searchKeyWord, "key-word", "k", "", helpDescriptions["domain-search-key-word"])
	create.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	info.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	remove.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	remove.Flags().StringVarP(&DomainId, "id", "i", "", helpDescriptions["domain-id"])
	nsRecords.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	check.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
}
