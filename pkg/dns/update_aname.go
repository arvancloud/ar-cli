package dns

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/antihax/optional"
	"github.com/masihyeganeh/ar-cli/pkg/api"
	"github.com/masihyeganeh/ar-cli/pkg/api/models"
	"github.com/masihyeganeh/ar-cli/pkg/utl"
	"github.com/spf13/cobra"
	"io"
)

// NewCmdDnsUpdateAName returns new cobra command to update DNS AName record
func NewCmdDnsUpdateAName(in io.Reader, out, errout io.Writer) *cobra.Command {
	var locationFlag string
	var hostHeaderFlag string
	var nameFlag string
	var ttlFlag int
	var cloudFlag bool
	var upstreamHttps string
	// Main command
	cmd := &cobra.Command{
		Use:   "aname domain id",
		Short: "update AName record #id for domain name (Example: example.com 1)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
		ValidArgs: []string{"domain", "id"},
		Args:      cobra.MinimumNArgs(2),
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)

			domain := args[0]
			id := args[1]

			record := models.DnsRecord{
				Type:  "aname",
				Cloud: cloudFlag,
			}

			if len(nameFlag) > 0 {
				record.Name = nameFlag
			}
			if len(locationFlag) > 0 {
				anameRecord := models.AnameRecord{
					Location: locationFlag,
				}
				if len(hostHeaderFlag) > 0 {
					anameRecord.HostHeader = hostHeaderFlag
				}

				record.Value = &models.OneOfDnsRecordValue{AnameRecord: anameRecord}
			}

			if ttlFlag == 120 || ttlFlag == 180 || ttlFlag == 300 || ttlFlag == 600 || ttlFlag == 900 || ttlFlag == 1800 || ttlFlag == 3600 || ttlFlag == 7200 || ttlFlag == 18000 || ttlFlag == 43200 || ttlFlag == 86400 || ttlFlag == 172800 || ttlFlag == 432000 {
				record.Ttl = int32(ttlFlag)
			}
			if upstreamHttps == "default" || upstreamHttps == "auto" || upstreamHttps == "http" || upstreamHttps == "https" {
				record.UpstreamHttps = upstreamHttps
			}
			// TODO:
			//record.IpFilterMode = &models.DnsRecordIpFilterMode{
			//	Count:     "",
			//	Order:     "",
			//	GeoFilter: "",
			//}

			options := &api.DNSApiDnsRecordsUpdateOpts{
				Body: optional.NewInterface(record),
			}
			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsUpdate(c.Context(), domain, id, options)
			utl.CheckApiErr(err)

			fmt.Fprintf(explainOut, "%s\n", res.Message)
		},
	}

	cmd.Flags().StringVarP(&locationFlag, "location", "o", "", "")
	cmd.Flags().StringVarP(&hostHeaderFlag, "host_header", "e", "", "")
	cmd.Flags().StringVarP(&nameFlag, "name", "n", "", "<= 250 characters")
	cmd.Flags().IntVarP(&ttlFlag, "ttl", "t", 0, "120 or 180 or 300 or 600 or 900 or 1800 or 3600 or 7200 or 18000 or 43200 or 86400 or 172800 or 432000")
	cmd.Flags().BoolVarP(&cloudFlag, "cloud", "l", false, "")
	cmd.Flags().StringVarP(&upstreamHttps, "upstream_https", "u", "default", `"default" or "auto" or "http" or "https"`)

	return cmd
}
