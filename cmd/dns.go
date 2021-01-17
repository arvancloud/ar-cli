package cmd

import (
	"encoding/json"
	"github.com/ebrahimahmadi/ar-cli/pkg/api"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
	"io/ioutil"
	"strconv"

	"github.com/spf13/cobra"
)

type DNSList struct {
	Data []DnsRecord
}

type DnsRecord struct {
	ID      string `json:"id"`
	DNSType string `json:"type"`
	Name    string `json:"name"`
	Value   struct {
		Host string `json:"host"`
	} `json:"value"`
	TTL           int    `json:"ttl"`
	Cloud         bool   `json:"cloud"`
	UpstreamHTTPS string `json:"upstream_https"`
	IPFilter      struct {
		Count     string `json:"count"`
		Order     string `json:"order"`
		GEOFilter string `json:"geo_filter"`
	} `json:"ip_filter_mode"`
	CanDelete   bool `json:"can_delete"`
	IsProtected bool `json:"is_protected"`
}

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Interact with Arvan DNS",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var dnsList = &cobra.Command{
	Use:   "list",
	Short: "Get list of DNS records",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-records",
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

		var dnsRecords = new(DNSList)
		_ = json.Unmarshal(responseData, &dnsRecords)

		table := newTable([]string{"Type", "Name", "Host", "TTL", "Cloud", "Upstream HTTPS", "Protected?", "IP GEO Filter", "IP Filter Order", "IP Filter Count"})

		for _, foundDomain := range dnsRecords.Data {
			record := []string{
				foundDomain.DNSType,
				foundDomain.Name,
				foundDomain.Value.Host,
				strconv.Itoa(foundDomain.TTL),
				strconv.FormatBool(foundDomain.Cloud),
				foundDomain.UpstreamHTTPS,
				strconv.FormatBool(foundDomain.IsProtected),
				foundDomain.IPFilter.GEOFilter,
				foundDomain.IPFilter.Order,
				foundDomain.IPFilter.Count,
			}
			table.Append(record)
		}

		table.Render()
	},
}
func init() {
	rootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(dnsList)
	dnsList.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	dnsList.MarkFlagRequired("name")
}
