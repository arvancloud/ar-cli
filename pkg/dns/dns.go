package dns

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/antihax/optional"
	"github.com/masihyeganeh/ar-cli/pkg/api/models"
	"github.com/spf13/cobra"
	"io"
	"os"

	"github.com/masihyeganeh/ar-cli/pkg/api"
	"github.com/masihyeganeh/ar-cli/pkg/utl"
)

// NewCmdDns returns new cobra command to manage DNSes
func NewCmdDns(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Manage DNS configurations",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
	}

	listCommand := NewCmdDnsList(in, out, errout)
	cmd.AddCommand(listCommand)

	createCommand := NewCmdDnsCreate(in, out, errout)
	cmd.AddCommand(createCommand)

	updateCommand := NewCmdDnsUpdate(in, out, errout)
	cmd.AddCommand(updateCommand)

	removeCommand := NewCmdDnsRemove(in, out, errout)
	cmd.AddCommand(removeCommand)

	toggleCloudCommand := NewCmdDnsToggleCloud(in, out, errout)
	cmd.AddCommand(toggleCloudCommand)

	importCommand := NewCmdDnsImport(in, out, errout)
	cmd.AddCommand(importCommand)

	return cmd
}

// NewCmdDnsList returns new cobra command to list DNS records from Arvan cloud
func NewCmdDnsList(in io.Reader, out, errout io.Writer) *cobra.Command {
	var searchFlag string
	var pageFlag int
	var perPageFlag int
	// Main command
	cmd := &cobra.Command{
		Use:   "list domain",
		Short: "List DNS configuration of domain name (Example: example.com)",
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

			options := &api.DNSApiDnsRecordsListOpts{}
			if len(searchFlag) > 0 {
				options.Search = optional.NewString(searchFlag)
			}
			if pageFlag > 1 {
				options.Page = optional.NewInt32(int32(pageFlag))
			}
			// It's a good idea to limit maximum perPage
			if pageFlag > 0 && perPageFlag < 100 {
				options.PerPage = optional.NewInt32(int32(perPageFlag))
			}

			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsList(c.Context(), domain, options)
			utl.CheckErr(err)

			fmt.Fprintf(explainOut, "%v\n", res.Data)
		},
	}

	cmd.Flags().StringVarP(&searchFlag, "search", "s", "", "Search term")
	cmd.Flags().IntVarP(&pageFlag, "page", "p", 1, "Set the desired page number")
	cmd.Flags().IntVarP(&perPageFlag, "per_page", "c", 0, "Set how many items returned per page")

	return cmd
}

// NewCmdDnsCreate returns new cobra command to create a DNS record
func NewCmdDnsCreate(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create DNS record for domain name (Example: example.com)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
	}

	aRecordCommand := NewCmdDnsCreateA(in, out, errout)
	cmd.AddCommand(aRecordCommand)

	aaaaRecordCommand := NewCmdDnsCreateAAAA(in, out, errout)
	cmd.AddCommand(aaaaRecordCommand)

	nsRecordCommand := NewCmdDnsCreateNS(in, out, errout)
	cmd.AddCommand(nsRecordCommand)

	txtRecordCommand := NewCmdDnsCreateTXT(in, out, errout)
	cmd.AddCommand(txtRecordCommand)

	cnameRecordCommand := NewCmdDnsCreateCName(in, out, errout)
	cmd.AddCommand(cnameRecordCommand)

	mxRecordCommand := NewCmdDnsCreateMX(in, out, errout)
	cmd.AddCommand(mxRecordCommand)

	srvRecordCommand := NewCmdDnsCreateSRV(in, out, errout)
	cmd.AddCommand(srvRecordCommand)

	spfRecordCommand := NewCmdDnsCreateSPF(in, out, errout)
	cmd.AddCommand(spfRecordCommand)

	dkimRecordCommand := NewCmdDnsCreateDKIM(in, out, errout)
	cmd.AddCommand(dkimRecordCommand)

	anameRecordCommand := NewCmdDnsCreateAName(in, out, errout)
	cmd.AddCommand(anameRecordCommand)

	ptrRecordCommand := NewCmdDnsCreatePTR(in, out, errout)
	cmd.AddCommand(ptrRecordCommand)

	return cmd
}

// NewCmdDnsUpdate returns new cobra command to update a DNS record
func NewCmdDnsUpdate(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update DNS record for domain name (Example: example.com)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
	}

	aRecordCommand := NewCmdDnsUpdateA(in, out, errout)
	cmd.AddCommand(aRecordCommand)

	aaaaRecordCommand := NewCmdDnsUpdateAAAA(in, out, errout)
	cmd.AddCommand(aaaaRecordCommand)

	nsRecordCommand := NewCmdDnsUpdateNS(in, out, errout)
	cmd.AddCommand(nsRecordCommand)

	txtRecordCommand := NewCmdDnsUpdateTXT(in, out, errout)
	cmd.AddCommand(txtRecordCommand)

	cnameRecordCommand := NewCmdDnsUpdateCName(in, out, errout)
	cmd.AddCommand(cnameRecordCommand)

	mxRecordCommand := NewCmdDnsUpdateMX(in, out, errout)
	cmd.AddCommand(mxRecordCommand)

	srvRecordCommand := NewCmdDnsUpdateSRV(in, out, errout)
	cmd.AddCommand(srvRecordCommand)

	spfRecordCommand := NewCmdDnsUpdateSPF(in, out, errout)
	cmd.AddCommand(spfRecordCommand)

	dkimRecordCommand := NewCmdDnsUpdateDKIM(in, out, errout)
	cmd.AddCommand(dkimRecordCommand)

	anameRecordCommand := NewCmdDnsUpdateAName(in, out, errout)
	cmd.AddCommand(anameRecordCommand)

	ptrRecordCommand := NewCmdDnsUpdatePTR(in, out, errout)
	cmd.AddCommand(ptrRecordCommand)

	return cmd
}

// NewCmdDnsRemove returns new cobra command to remove a DNS record from Arvan cloud
func NewCmdDnsRemove(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "remove domain id",
		Short: "Remove DNS record #id from domain name (Example: example.com 1)",
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

			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsRemove(c.Context(), domain, id)
			utl.CheckErr(err)

			fmt.Fprintf(explainOut, "%v\n", res.Data)
		},
	}

	return cmd
}

// NewCmdDnsToggleCloud returns new cobra command to toggle cloud of a DNS record from Arvan cloud
func NewCmdDnsToggleCloud(in io.Reader, out, errout io.Writer) *cobra.Command {
	var cloudFlag bool
	// Main command
	cmd := &cobra.Command{
		Use:   "toggle domain id",
		Short: "Toggle cloud of DNS record #id from domain name (Example: example.com 1)",
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

			options := &api.DNSApiDnsRecordsCloudOpts{
				Body: optional.NewInterface(models.CloudToggleBody{Cloud: cloudFlag}),
			}

			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsCloud(c.Context(), domain, id, options)
			utl.CheckErr(err)

			fmt.Fprintf(explainOut, "%v\n", res.Data)
		},
	}

	cmd.Flags().BoolVarP(&cloudFlag, "cloud", "l", false, "To proxy or not proxy, that's the question!")

	return cmd
}

// NewCmdDnsImport returns new cobra command to import DNS records from BIND file
func NewCmdDnsImport(in io.Reader, out, errout io.Writer) *cobra.Command {
	// Main command
	cmd := &cobra.Command{
		Use:   "import domain f_zone_file",
		Short: "import DNS configuration from f_zone_file for domain name (Example: example.com)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
		ValidArgs: []string{"domain", "f_zone_file"},
		Args:      cobra.MinimumNArgs(2),
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)

			domain := args[0]
			fZoneFile := args[1]

			file, err := os.Open(fZoneFile)
			utl.CheckErr(err)

			options := &api.DNSApiDnsRecordsImportOpts{
				FZoneFile: file,
			}

			defer file.Close()

			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsImport(c.Context(), domain, options)
			utl.CheckErr(err)

			fmt.Fprintf(explainOut, "%v\n", res.Data)
		},
	}

	return cmd
}
