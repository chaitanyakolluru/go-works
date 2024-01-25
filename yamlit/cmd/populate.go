/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"yamlit/cmd/populateyaml"

	"github.com/spf13/cobra"
)

var pattern, vertical, appname, componentname, storageaccount, container, key, accesskey string

// populateCmd represents the populate command
var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populate yaml that generates variables-in-mouldes.yaml file.",
	Long: `Populate Yaml:
This subcommand takes in the filled in service-offerings.yaml file from the discover services subcommand step, and creates a new file called "variables-in-modules.yaml"
in the current working directory with all variables each of the modules need. You need to provide values to each variable within the modules and can proceed to invoke the 
next two steps.
	
Example:

$ yamlit populate

Creates:
If Service_Principal_Authentication is not specified in service-offerings.yaml:
$ cat variables-in-modules.yaml (if appservice is selected in previous step)
appservice:
  appservice: FILL OUT
  appservice_plan_kind: FILL OUT
  appservice_plan_name: FILL OUT
  appservice_plan_reserved: FILL OUT
  appservice_plan_size: FILL OUT
  appservice_plan_tier: FILL OUT
  appservice_slot: FILL OUT
  custom_tags: FILL OUT
  dynamic_app_settings: FILL OUT
  existing_appservice_plan_name: FILL OUT
  identity_type: FILL OUT
  keyvault_name: FILL OUT
  resource_group_name: FILL OUT
  version: FILL OUT

If Service_Principal_Authentication is specified in service-offerings.yaml: (Creates entries to populate with service principal auth details)
$ cat variables-in-modules.yaml (if keyvault is selected in previous step)
client_id: FILL OUT
client_secret: FILL OUT
keyvault:
  access_policy: FILL OUT
  enabled_for_deployment: FILL OUT
  enabled_for_disk_encryption: FILL OUT
  enabled_for_template_deployment: FILL OUT
  keyvault_name: FILL OUT
  keyvault_sku: FILL OUT
  purge_protection_enabled: FILL OUT
  resource_group_name: FILL OUT
  secrets: FILL OUT
  soft_delete_enabled: FILL OUT
  version: FILL OUT
subscription_id: FILL OUT
tenant_id: FILL OUT

Please fill out the values for resource attribute in yaml format. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		if pattern != "" {
			if vertical == "" || appname == "" || componentname == "" || storageaccount == "" || container == "" || key == "" {
				cmd.Usage()
				os.Exit(1)
			}
		}

		populateyaml.PopulateYaml(cwd, pattern, vertical, appname, componentname, storageaccount, container, key, accesskey, username, token)
	},
}

func init() {
	rootCmd.AddCommand(populateCmd)
	populateCmd.Flags().StringVarP(&pattern, "pattern", "", "", "Flag that lets you create variables-in-modules.yaml file, partly filled in with attributes from the pattern specified")
	populateCmd.Flags().StringVarP(&vertical, "vertical", "", "", "Flag that lets you provide vertical to the populate command")
	populateCmd.Flags().StringVarP(&appname, "appname", "", "", "Flag that lets you provide application short name to the populate command")
	populateCmd.Flags().StringVarP(&componentname, "componentname", "", "", "Flag that lets you provide component name to the populate command")
	populateCmd.Flags().StringVarP(&storageaccount, "storageaccount", "", "", "Flag that lets you provide storage account name to the populate command for remote statefile mangement")
	populateCmd.Flags().StringVarP(&container, "container", "", "", "Flag that lets you provide container name under the storage account to the populate command for remote statefile management")
	populateCmd.Flags().StringVarP(&key, "key", "", "", "Flag that lets you provide key within the storage account to the populate command for remote statefile management")
	populateCmd.Flags().StringVarP(&accesskey, "accesskey", "", "", "Flag that lets you provide accesskey of the storage account to the populate command for remote statefile management")

}
