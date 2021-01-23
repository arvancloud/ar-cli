package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ebrahimahmadi/ar-cli/internals/pkg/utils"
	"github.com/ebrahimahmadi/ar-cli/pkg/api"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
	"io/ioutil"

	"github.com/spf13/cobra"
)

type CdnList struct {
	Data []CDNData
}

type CDNInfo struct {
	Data CDNData
}

type CDNData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortDesc string `json:"short_description"`
	Vendor    string `json:"vendor"`
	Status    string `json:"status"`
}

var cdnId string
var eventToTrigger string
var CdnTriggerEvents = []string{
	"before-new-install",
	"new-install",
}

var cdnAppCmd = &cobra.Command{
	Use:   "cdn",
	Short: "Interact with cdn app service",
	Long:  helpDescriptions["cdnapp-command"],
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cdnApp called")
	},
}

var cdnAppList = &cobra.Command{
	Use:   "apps",
	Short: "Get list of all available cdn-apps",
	Long:  helpDescriptions["cdnapp-list"],
	Run: func(cmd *cobra.Command, args []string) {
		request := api.RequestBag{
			URL:    Config.GetUrl() + "/apps",
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

		var cdns = new(CdnList)
		_ = json.Unmarshal(responseData, &cdns)

		table := utils.NewTable([]string{"ID", "Name", "Vendor", "Status"})

		for _, availableCDN := range cdns.Data {
			record := []string{
				availableCDN.Id,
				availableCDN.Name,
				availableCDN.Vendor,
				availableCDN.Status,
			}
			table.Append(record)
		}

		table.Render()
	},
}

var cdnAppInfo = &cobra.Command{
	Use:   "info",
	Short: "Find Cdn app by id",
	Run: func(cmd *cobra.Command, args []string) {
		request := api.RequestBag{
			URL:    Config.GetUrl() + "/apps/" + cdnId,
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

		var cdn = new(CDNInfo)
		_ = json.Unmarshal(responseData, &cdn)

		table := utils.NewTable([]string{"ID", "Name", "Vendor", "Status"})

		record := []string{
			cdn.Data.Id,
			cdn.Data.Name,
			cdn.Data.Vendor,
			cdn.Data.Status,
		}

		table.Append(record)

		table.Render()
	},
}

var installedApp = &cobra.Command{
	Use:   "installed-apps",
	Short: "Get list of all applications installed on a domain",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/apps",
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

		var cdn = new(CDNInfo)
		_ = json.Unmarshal(responseData, &cdn)

		table := utils.NewTable([]string{"ID", "Name", "Vendor", "Status"})

		record := []string{
			cdn.Data.Id,
			cdn.Data.Name,
			cdn.Data.Vendor,
			cdn.Data.Status,
		}

		table.Append(record)
		table.Render()
	},
}


var installApp = &cobra.Command{
	Use:   "install",
	Short: "Install the application on the domain",
	Run: func(cmd *cobra.Command, args []string) {
		_, domainValidationErr := validator.IsDomain(DomainName)

		if domainValidationErr != nil {
			err := helpers.ToBeColored{Expression: domainValidationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/apps/" + cdnId,
			Method: "POST",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		info := helpers.ToBeColored{Expression: "Application successfully installed on " + DomainName}
		info.StdoutNotice()
	},
}


var uninstallApp = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall application from the domain",
	Run: func(cmd *cobra.Command, args []string) {
		_, domainValidationErr := validator.IsDomain(DomainName)

		if domainValidationErr != nil {
			err := helpers.ToBeColored{Expression: domainValidationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/apps/" + cdnId,
			Method: "DELETE",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		info := helpers.ToBeColored{Expression: "Application successfully uninstalled from " + DomainName}
		info.StdoutNotice()
	},
}


var triggerWebHook = &cobra.Command{
	Use:   "trigger",
	Short: "trigger webhook event",
	Run: func(cmd *cobra.Command, args []string) {
		_, domainValidationErr := validator.IsDomain(DomainName)
		_, eventValidationErr := validator.HasString(eventToTrigger, CdnTriggerEvents)

		if domainValidationErr != nil {
			err := helpers.ToBeColored{Expression: domainValidationErr.Error()}
			err.StdoutError().StopExecution()
		}

		if eventValidationErr != nil {
			err := helpers.ToBeColored{Expression: eventValidationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			BodyPayload: map[string]interface{}{
				"event": eventToTrigger,
			},
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/apps/" + cdnId + "/actions/trigger_webhook",
			Method: "POST",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		info := helpers.ToBeColored{Expression: "Application successfully uninstalled from " + DomainName}
		info.StdoutNotice()
	},
}

func init() {
	rootCmd.AddCommand(cdnAppCmd)
	cdnAppCmd.AddCommand(cdnAppList)
	cdnAppCmd.AddCommand(cdnAppInfo)
	cdnAppCmd.AddCommand(installedApp)
	cdnAppCmd.AddCommand(installApp)
	cdnAppCmd.AddCommand(uninstallApp)
	cdnAppCmd.AddCommand(triggerWebHook)

	cdnAppInfo.Flags().StringVarP(&cdnId, "id", "i", "", helpDescriptions["cdnapp-id"])
	cdnAppInfo.MarkFlagRequired("id")

	installedApp.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	installedApp.MarkFlagRequired("name")

	installApp.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	installApp.Flags().StringVarP(&cdnId, "id", "i", "", helpDescriptions["cdnapp-id"])
	installApp.MarkFlagRequired("name")
	installApp.MarkFlagRequired("id")

	uninstallApp.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	uninstallApp.Flags().StringVarP(&cdnId, "id", "i", "", helpDescriptions["cdnapp-id"])
	uninstallApp.MarkFlagRequired("name")
	uninstallApp.MarkFlagRequired("id")

	triggerWebHook.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	triggerWebHook.Flags().StringVarP(&cdnId, "id", "i", "", helpDescriptions["cdnapp-id"])
	triggerWebHook.Flags().StringVarP(&eventToTrigger, "event", "e", "", helpDescriptions["cdnapp-id"])
	triggerWebHook.MarkFlagRequired("name")
	triggerWebHook.MarkFlagRequired("id")
	triggerWebHook.MarkFlagRequired("event")
}
