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
	"yamlit/cmd/generatemaintf"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "This option take in the filled in variables-in-modules.yaml and generates a main.tf file under cwd/.temp.",
	Long: `This option take in the filled in variables-in-modules.yaml and generates a main.tf file under cwd/.temp.
	
Example:
$ yamlit generate
	
Input example:
$ cat variables-in-modules.yaml (without Service_Principal_Authentication requested in previous steps)
appservice:
  appservice:
    appservice-1:
	  appservice_name: appserv-demo-1
	appservice-2:
	  appservice_name: appserv-demo-2
	appservice-3:
	  appservice_name: appserv-demo-3
  appservice_plan_kind: Linux
  appservice_plan_name: appserv-demo-plan
  appservice_plan_reserved: true
  appservice_plan_size: S1
  appservice_plan_tier: Standard
  resource_group_name: aa-demo-rg
  version: appservice-v2.2.1

$ cat variables-in-modules.yaml (with Service_Principal_Authentication requested in previous steps)
subscription_id: "subscription_id"
tenant_id: "tenant_id"
client_id: "client_id"
client_secret: "client_secret"
appservice:
  appservice:
    appservice-1:
	  appservice_name: appserv-demo-1
	appservice-2:
	  appservice_name: appserv-demo-2
	appservice-3:
	  appservice_name: appserv-demo-3
  appservice_plan_kind: Linux
  appservice_plan_name: appserv-demo-plan
  appservice_plan_reserved: true
  appservice_plan_size: S1
  appservice_plan_tier: Standard
  resource_group_name: aa-demo-rg
version: appservice-v2.2.1
	  `,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		if inputyaml == "" {

			if environment != "" {
				if location == "" {
					cmd.Usage()
					os.Exit(1)
				}
			}

			if location != "" {
				if environment == "" {
					cmd.Usage()
					os.Exit(1)
				}

			}
		}
		generatemaintf.GenerateMaintf(cwd, inputyaml, environment, location, username, token)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&inputyaml, "inputyaml", "", "", "Flag that lets you provide the name of the yaml file that needs to be processed")
	generateCmd.Flags().StringVarP(&environment, "environment", "", "", "Flag that lets you provide the environment of application infrastructure")
	generateCmd.Flags().StringVarP(&location, "location", "", "", "Flag that lets you provide the location of application infrastructure")
}
