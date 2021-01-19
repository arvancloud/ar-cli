package cli

import (
	"fmt"
	"github.com/masihyeganeh/ar-cli/pkg/dns"
	"github.com/masihyeganeh/ar-cli/pkg/dnssec"
	"github.com/masihyeganeh/ar-cli/pkg/utl"
	"github.com/spf13/cobra"
	"os"

	"github.com/masihyeganeh/ar-cli/pkg/config"
	"github.com/masihyeganeh/ar-cli/pkg/login"
)

var (
	cliName = "ar-cli"
	cliLong = `
    Arvan CDN Services
    This client helps you manage CDN in Arvan Cloud Services`

	cliExplain = `
    To use manage DNS of your domains run ar-cli command with your service name:
        ar-cli dns --help
    This will show manual for managing Arvan DNS services.
    To see the full list of commands supported, run 'ar-cli --help'.`
)

// NewCommandCLI return new cobra cli
func NewCommandCLI() *cobra.Command {
	// Load ConfigInfo from default path if exists
	_, _ = config.LoadConfigFile()

	in, out, errout := os.Stdin, os.Stdout, os.Stderr
	// Main command
	cmd := &cobra.Command{
		Use:   cliName,
		Short: "Command line tools for managing Arvan CDN services",
		Long:  cliLong,
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)
			fmt.Fprintf(explainOut, "%s\n\n%s\n", cliLong, cliExplain)
		},
	}

	optionsCommand := newCmdOptions()
	cmd.AddCommand(optionsCommand)

	loginCommand := login.NewCmdLogin(in, out, errout)
	cmd.AddCommand(loginCommand)

	dnsCommand := dns.NewCmdDns(in, out, errout)
	cmd.AddCommand(dnsCommand)

	dnsSecCommand := dnssec.NewCmdDnsSec(in, out, errout)
	cmd.AddCommand(dnsSecCommand)

	return cmd
}

// newCmdOptions implements the OpenShift cli options command
func newCmdOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use: "options",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return cmd
}
