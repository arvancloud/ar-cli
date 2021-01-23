package cmd

import (
	"encoding/json"
	"github.com/ebrahimahmadi/ar-cli/internals/utils"
	"github.com/ebrahimahmadi/ar-cli/pkg/api"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
)

type SecurityInfo struct {
	Data struct {
		Plan             string `json:"plan"`
		WAFStatus        string `json:"waf_status"`
		DDOSType         string `json:"ddos_type"`
		FWRules          int    `json:"firewall_rules"`
		LimitationStatus bool   `json:"limitation_status"`
	} `json:"data"`
}

var cloudSecurityPlan string
var cloudSecurityAvailablePlans = []string{
	"bronze",
	"silver",
	"gold",
	"platinum",
}

var cloudSecurityCmd = &cobra.Command{
	Use:   "cloud-security",
	Short: "Check cloud security status or update your plan",
	Long:  helpDescriptions["cs-command"],
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var csServicesStatus = &cobra.Command{
	Use:   "info",
	Short: "Get an overview of cloud security services status",
	Long:  helpDescriptions["cs-info"],
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/security-service/info",
			Method: "GET",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		responseData, _ := ioutil.ReadAll(res.Body)

		var securityInfo = new(SecurityInfo)
		_ = json.Unmarshal(responseData, &securityInfo)

		api.HandleResponseErr(res)

		table := utils.NewTable([]string{"Subscribed Plan", "WAF", "DDOS Type", "Firewall Rules", "Limitation Status"})

		record := []string{
			securityInfo.Data.Plan,
			securityInfo.Data.WAFStatus,
			securityInfo.Data.DDOSType,
			strconv.Itoa(securityInfo.Data.FWRules),
			strconv.FormatBool(securityInfo.Data.LimitationStatus),
		}

		table.Append(record)
		table.Render()
	},
}

var csUpdatePlan = &cobra.Command{
	Use:   "update",
	Short: "Update you subscription",
	Long:  helpDescriptions["cs-update-plan"],
	Run: func(cmd *cobra.Command, args []string) {
		_, enumValidation := validator.HasString(cloudSecurityPlan, cloudSecurityAvailablePlans)

		if enumValidation != nil {
			err := helpers.ToBeColored{Expression: enumValidation.Error()}
			err.StdoutError().StopExecution()
		}

		_, domainNameValidation := validator.IsDomain(DomainName)

		if domainNameValidation != nil {
			err := helpers.ToBeColored{Expression: domainNameValidation.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			BodyPayload: map[string]interface{}{"plan": cloudSecurityPlan},
			URL:         Config.GetUrl() + "/domains/" + DomainName + "/security-service/plan",
			Method:      "PUT",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		colored := helpers.ToBeColored{Expression: "Cloud Security Updated"}
		colored.StdoutNotice()
	},
}

func init() {
	rootCmd.AddCommand(cloudSecurityCmd)

	cloudSecurityCmd.AddCommand(csServicesStatus)
	cloudSecurityCmd.AddCommand(csUpdatePlan)

	csServicesStatus.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	csServicesStatus.MarkFlagRequired("name")

	csUpdatePlan.Flags().StringVarP(&cloudSecurityPlan, "plan", "p", "", helpDescriptions["cs-plan"])
	csUpdatePlan.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	csUpdatePlan.MarkFlagRequired("name")
	csUpdatePlan.MarkFlagRequired("plan")
}
