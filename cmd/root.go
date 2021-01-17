package cmd

import (
	"fmt"
	"github.com/ebrahimahmadi/ar-cli/pkg/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var DomainName string
var DomainId string
var Config = config.GetConfigInfo()

var helpDescriptions = map[string]string{
	"domain-command":   "Create, Search, Delete, Get, Health check and get Ns records ",
	"domain-search":          "Leaving the 'search' flag empty, will return all domains. Otherwise, it will filter domains containing the search keyword.",
	"domain-search-key-word": "Search Item",
	"domain-create":          "Create new domain",
	"domain-info":            "Get information of the domain",
	"domain-name":     "The host name. like: example.com",
	"domain-id":       "The domain UUID. like: 3541b0ce-e8a6-42f0-b65a-f03a7c387486",
	"domain-remove":          "Remove the domain",
	"domain-list-ns-records": "Get list of domain's root NS records and expected values",
	"domain-check":           "Check NS to find whether domain is activated",

	"cs-command": `Get an overview of cloud security services status. Or update your security plans`,
	"cs-info":  "Get an overview of cloud security services status",
	"cs-update-plan": "Update your cloud security plan",
	"cs-plan": "The plan you are wiling to subscribe to. The value should be one of: bronze, silver, gold, platinum",

	"cdnapp-command": "Find, List, Install and Uninstall the application from domain",
	"cdnapp-list": "List all available CDNs",
	"cdnapp-info": "get cdn using the given cdn id",
	"cdnapp-id": "The cdn app UUID you are wiling to receive id",
	"cdnapp-event": "Event that you want to trigger. The value should be on of: before-new-install, new-install",

	"dns-record-id": "ID of the DNS record",
	"dns-record-cloud": "If is true the cloud status will be served over cloud, otherwise no!",
}

func newTable(tableHeaders []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	center := 1

	table.SetAlignment(center)
	table.SetHeader(tableHeaders)
	table.SetRowLine(true)
	table.SetRowSeparator("~")

	return table
}

var rootCmd = &cobra.Command{
	Use:   "ar-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ar-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ar-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ar-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
