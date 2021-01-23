package cmd

import (
	"encoding/json"
	"github.com/ebrahimahmadi/ar-cli/internals/pkg/factories/dns_records"
	"github.com/ebrahimahmadi/ar-cli/internals/pkg/utils"
	"github.com/ebrahimahmadi/ar-cli/pkg/api"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
	"io/ioutil"
	"strconv"

	"github.com/spf13/cobra"
)

type RecordList struct {
	Data []DnsRecord `json:"data"`
}

type RecordInfo struct {
	Data DnsRecord `json:"data"`
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

var ipV4 string
var port int
var country string
var weight int
var ttl int
var recordName string
var upstreamHttps string
var ipFilterCount string
var ipFilterOrder string
var ipGeoFilter string
var isRecordServedOnCloud bool
var dnsRecordId string

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

		var dnsRecords = new(RecordList)
		_ = json.Unmarshal(responseData, &dnsRecords)

		table := utils.NewTable([]string{"ID", "Type", "Name", "Host", "TTL", "Cloud", "Upstream HTTPS", "Protected?", "IP GEO Filter", "IP Filter Order", "IP Filter Count"})

		for _, foundDomain := range dnsRecords.Data {
			record := []string{
				foundDomain.ID,
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

var dnsGet = &cobra.Command{
	Use:   "info",
	Short: "Get information of a record",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-records/" + dnsRecordId,
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

		var dnsRecord = new(RecordInfo)
		_ = json.Unmarshal(responseData, &dnsRecord)

		table := utils.NewTable([]string{"ID", "Type", "Name", "Host", "TTL", "Cloud", "Upstream HTTPS", "Protected?", "IP GEO Filter", "IP Filter Order", "IP Filter Count"})

		record := []string{
			dnsRecord.Data.ID,
			dnsRecord.Data.DNSType,
			dnsRecord.Data.Name,
			dnsRecord.Data.Value.Host,
			strconv.Itoa(dnsRecord.Data.TTL),
			strconv.FormatBool(dnsRecord.Data.Cloud),
			dnsRecord.Data.UpstreamHTTPS,
			strconv.FormatBool(dnsRecord.Data.IsProtected),
			dnsRecord.Data.IPFilter.GEOFilter,
			dnsRecord.Data.IPFilter.Order,
			dnsRecord.Data.IPFilter.Count,
		}
		table.Append(record)

		table.Render()
	},
}

var dnsRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove a DNS record",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-records/" + dnsRecordId,
			Method: "DELETE",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		notice := helpers.ToBeColored{Expression: "Removed Successfully"}
		notice.StdoutError()
	},
}

var dnsToggle = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle Cloud Status",
	Long:  "Toggle cloud status (To proxy or not proxy, that's the question!)",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			BodyPayload: map[string]interface{}{
				"cloud": isRecordServedOnCloud,
			},
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/dns-records/" + dnsRecordId + "/cloud",
			Method: "PUT",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		notice := helpers.ToBeColored{Expression: "Toggled Successfully"}
		notice.StdoutNotice()
	},
}

var dnsCreate = &cobra.Command{
	Use:   "create",
	Short: "Toggle Cloud Status",
	Long:  "Toggle cloud status (To proxy or not proxy, that's the question!)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var aRecord = &cobra.Command{
	Use:   "a-record",
	Short: "Create A type DNS record ",
	Run: func(cmd *cobra.Command, args []string) {

		if validDomain, validationErr := validator.IsDomain(DomainName); !validDomain {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		if validIp, ipValidationErr := validator.IsValidIp(ipV4); !validIp {
			err := helpers.ToBeColored{Expression: ipValidationErr.Error()}
			err.StdoutError().StopExecution()
		}

		record := dns_records.ARecord{
			TTL:           ttl,
			Name:          recordName,
			Cloud:         isRecordServedOnCloud,
			UpstreamHttps: upstreamHttps,
			Value: []dns_records.Value{
				{
					IP: ipV4,
					Country: country,
					Port: port,
					Weight: weight,
				},
			},
			IpFilterMode: dns_records.IpFilterMode{
				Count:     ipFilterCount,
				Order:     ipFilterOrder,
				GeoFilter: ipGeoFilter,
			},
		}

		body, err := record.Build()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		req := api.RequestBag{
			FormattedBodyPayload: body,
			URL:         Config.GetUrl() + "/domains/" + DomainName + "/dns-records",
			Method:      "POST",
		}

		res, err := req.Do()

		api.HandleResponseErr(res)

		notice := helpers.ToBeColored{Expression: "Created Successfully"}
		notice.StdoutNotice()
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(dnsList)
	dnsCmd.AddCommand(dnsGet)
	dnsCmd.AddCommand(dnsRemove)
	dnsCmd.AddCommand(dnsToggle)
	dnsCmd.AddCommand(dnsCreate)
	dnsCreate.AddCommand(aRecord)

	dnsList.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	dnsList.MarkFlagRequired("name")

	dnsGet.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	dnsGet.Flags().StringVarP(&dnsRecordId, "record-id", "r", "", helpDescriptions["dns-record-id"])
	dnsGet.MarkFlagRequired("name")
	dnsGet.MarkFlagRequired("record-id")

	dnsRemove.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	dnsRemove.Flags().StringVarP(&dnsRecordId, "record-id", "r", "", helpDescriptions["dns-record-id"])
	dnsRemove.MarkFlagRequired("name")
	dnsRemove.MarkFlagRequired("record-id")

	dnsToggle.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	dnsToggle.Flags().StringVarP(&dnsRecordId, "record-id", "r", "", helpDescriptions["dns-record-id"])
	dnsToggle.Flags().BoolVarP(&isRecordServedOnCloud, "cloud", "c", true, helpDescriptions["dns-record-cloud"])
	dnsToggle.MarkFlagRequired("name")
	dnsToggle.MarkFlagRequired("record-id")

	aRecord.Flags().StringVarP(&DomainName, "domain", "d", "", helpDescriptions["domain-name"])
	aRecord.Flags().StringVarP(&recordName, "name", "n", "", helpDescriptions["dns-record-name"])
	aRecord.Flags().StringVarP(&ipV4, "ip", "i", "", helpDescriptions["dns-ip"])
	aRecord.Flags().IntVarP(&port, "port", "p", 0, helpDescriptions["dns-port"])
	aRecord.Flags().IntVarP(&weight, "weight", "w", 0, helpDescriptions["dns-weight"])
	aRecord.Flags().IntVarP(&ttl, "ttl", "t", 120, helpDescriptions["dns-ttl"])
	aRecord.Flags().BoolVarP(&isRecordServedOnCloud, "cloud", "l", false, helpDescriptions["dns-record-cloud"])
	aRecord.Flags().StringVarP(&country, "country", "c", "", helpDescriptions["dns-country"])
	aRecord.Flags().StringVarP(&ipFilterCount, "ip-filter-count", "f", "", helpDescriptions["dns-filter-count"])
	aRecord.Flags().StringVarP(&ipFilterOrder, "ip-filter-order", "o", "", helpDescriptions["dns-filter-order"])
	aRecord.Flags().StringVarP(&ipGeoFilter, "ip-geo-filter", "g", "", helpDescriptions["dns-geo-filter"])
	aRecord.Flags().StringVarP(&upstreamHttps, "upstream-https", "u", "default", helpDescriptions["dns-upstream-httpa"])
	aRecord.MarkFlagRequired("ip")
	aRecord.MarkFlagRequired("domain-name")
	aRecord.MarkFlagRequired("upstream-https")
	aRecord.MarkFlagRequired("name")
	aRecord.MarkFlagRequired("ttl")
	aRecord.MarkFlagRequired("ip-filter-count")
	aRecord.MarkFlagRequired("ip-filter-order")
	aRecord.MarkFlagRequired("ip-geo-filter")

}
