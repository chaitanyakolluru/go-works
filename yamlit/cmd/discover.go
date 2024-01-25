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
	"yamlit/cmd/discoverservices"

	"github.com/spf13/cobra"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover services offered within AA/Terraform.",
	Long: `Discover Services:
This subcommand procures the modules currently supported by AA/Terraform and lists them in a yaml file named service-offerings.yaml in the current working directory. 
You can suggest the resources you want to create by providing a value of 1 as value to the resource's key in the yaml file.
	
Example:
$ yamlit discover

Creates:
$ cat services-offered.yaml 
Please select the items needed by placing a 1 beside the item..

Service_Principal_Authentication:

aks:
apim: 
app-insights:
appservice:
azure_sql: 
cdn-profile: 
container-registry:
cosmosdb: 
eventhub: 
function-app:
keyvault: 1
log-analytics:
mysql_server: 
resource-group:
route-table:
snet:
storage_account: 
terraform-examples: 
traffic-manager:
vnet:`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		discoverservices.DiscoverServices(cwd, username, token)
	},
}

func init() {
	rootCmd.AddCommand(discoverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
