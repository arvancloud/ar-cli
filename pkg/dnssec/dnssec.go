package dnssec

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

// NewCmdDnsSec returns new cobra command to manage DNSSECs
func NewCmdDnsSec(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "dnssec",
		Short: "Manage DNSSEC configurations",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
	}

	listCommand := NewCmdDnsSecStatus(in, out, errout)
	cmd.AddCommand(listCommand)

	toggleCloudCommand := NewCmdDnsSecUpdate(in, out, errout)
	cmd.AddCommand(toggleCloudCommand)

	return cmd
}

// NewCmdDnsSecStatus returns new cobra command to get status of DNSSEC for a domain from Arvan cloud
func NewCmdDnsSecStatus(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "status domain",
		Short: "Get DNSSEC status of domain name (Example: example.com)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
		ValidArgs: []string{"domain"},
		Args:      cobra.MinimumNArgs(1),
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)

			domain := args[0]

			res, _, err := api.GetAPIClient().DNSSECApi.DnsRecordsDnssecGet(c.Context(), domain)
			utl.CheckApiErr(err)

			if res.Data.Enabled {
				fmt.Fprintf(explainOut, "DNSSEC is enabled for %s. DS : %s\n", domain, res.Data.Ds)
			} else {
				fmt.Fprintf(explainOut, "DNSSEC is disabled for %s\n", domain)
			}
		},
	}

	return cmd
}

// NewCmdDnsSecUpdate returns new cobra command to toggle DNSSEC for a domain from Arvan cloud
func NewCmdDnsSecUpdate(in io.Reader, out, errout io.Writer) *cobra.Command {
	var enable bool
	// Main command
	cmd := &cobra.Command{
		Use:   "update domain",
		Short: "Update DNSSEC of domain name (Example: example.com -e)",
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
		ValidArgs: []string{"domain"},
		Args:      cobra.MinimumNArgs(1),
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)

			domain := args[0]

			options := &api.DNSSECApiDnsRecordsDnssecActionsOpts{
				Body: optional.NewInterface(models.DnsSecUpdateBody{Enable: enable}),
			}

			res, _, err := api.GetAPIClient().DNSSECApi.DnsRecordsDnssecActions(c.Context(), domain, options)
			utl.CheckApiErr(err)

			if res.Data.Enabled {
				fmt.Fprintf(explainOut, "DNSSEC is enabled for %s. DS : %s\n", domain, res.Data.Ds)
			} else {
				fmt.Fprintf(explainOut, "DNSSEC is disabled for %s\n", domain)
			}
		},
	}

	cmd.Flags().BoolVarP(&enable, "enable", "e", false, "To enable DNSSEC")

	return cmd
}
