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
	"yamlit/cmd/terraformruns"
	"yamlit/cmd/testterraform"

	"github.com/spf13/cobra"
)

var testfile, testcondition string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test subcommand lets you provide a --testfile and a --testcondition file, run terraform apply and test each condition aginst the resulting terraform.tfstate file.",
	Long: `Test:
Test command lets you provide terraform code with flag --testfile and a condition file written in yaml using flag --testcondition. The testcondition yaml file contains 
nested map data in the following format:
	
resouce1:
resource2:
	property1: value1
	property2: value2
....
		
The above format dictates that resource 1 needs to be created and that resource 2, along with being created has property1 set to value1 and property2 set to value2.
	
With the --testfile and --testcondition provided, the subcommand then proceeds to use the --testfile and runs terraform init and apply over the terraform code. After creation,
it proceeds to process the end state as recorded in the terraform.tfstate file and uses that to test if the resources and resource properties defined within
the --testcondition file match with the resource state obtained from the terraform.tfstate file. 
	
Once test finishes, the subcommand returns the results of the test as shown below:

$ cat temp/demo2.tf
provider azurerm {
	features {}
}
	
module "storageaccount_demo" {
	source                       = "git@ghe.aa.com:AA/terraform.git//azure-modules/storage_account"
	storage_account_name         = "storageaccountnamedemo"
	resource_group_name          = "aa-aot-sais-nonprod-eastus-rg"
	account_tier                 = "Standard"
	account_kind                 = "StorageV2"
	account_replication_type     = "grs"
	access_tier                  = "Hot"
	network_rules_default_action = "Allow"
	network_rules_bypass         = ["AzureServices"]
	enable_https_traffic_only    = true
	blob_container = {
	   container1 = {
			container_name        = "democontainer1",
			storage_account_name  = "storageaccountnamedemo",
			container_access_type = "container"
		},
		container2 = {
			container_name        = "democontainer2",
			storage_account_name  = "storageaccountnamedemo",
			container_access_type = "container"
		},
	}
}
	
module "appservice_demo" {
	source                   = "git@ghe.aa.com:AA/terraform.git//azure-modules/appservice"
	appservice_plan_name     = "appserv-chai-plan"
	appservice_plan_tier     = "Standard"
	appservice_plan_size     = "S1"
	appservice_plan_kind     = "Linux"
	resource_group_name      = "aa-aot-sais-nonprod-eastus-rg"
	appservice_plan_reserved = true
	appservice = {
		"appservice-1": {
			"appservice_name": "appserv-chai-1"
		},
		"appservice-2": {
			"appservice_name": "appserv-chai-2"
		}
	}
}

---
	
$cat temp/demo2-testcondition.yaml 
storageaccountnamedemo:
	location: eastus
	resource_group_name: aa-aot-sais-nonprod-eastus-rg
	account_tier: Standard
	account_kind: StorageV2
	account_replication_type: GRS
	access_tier: Hot
	default_action: Allow
	bypass: AzureServices
	enable_https_traffic_only: true
	democontainer1:
		container_access_type: "container"
	democontainer2:
	  container_access_type: "container"
	appserv-chai-plan:
		kind: linux
		size: S1
		tier: Standard
	appserv-chai-1:
		min_tls_version: "1.2"
	appserv-chai-2:
	
---

$ yamlit test --testfile temp/demo2.tf --testcondition temp/demo2-testcondition.yaml
storageaccountnamedemo:enable_https_traffic_only:true :: TRUE
storageaccountnamedemo:bypass:AzureServices :: TRUE
storageaccountnamedemo:location:eastus :: TRUE
storageaccountnamedemo:resource_group_name:aa-aot-sais-nonprod-eastus-rg :: TRUE
storageaccountnamedemo:account_tier:Standard :: TRUE
storageaccountnamedemo:account_kind:StorageV2 :: TRUE
storageaccountnamedemo:account_replication_type:GRS :: TRUE
storageaccountnamedemo:access_tier:Hot :: TRUE
storageaccountnamedemo:default_action:Allow :: TRUE
democontainer1:container_access_type:container :: TRUE
democontainer2:container_access_type:container :: TRUE
appserv-chai-plan:kind:linux :: TRUE
appserv-chai-plan:size:S1 :: TRUE
appserv-chai-plan:tier:Standard :: TRUE
appserv-chai-1:min_tls_version:1.2 :: TRUE
appserv-chai-2 :: TRUE
=======================================================================================
==============================TESTS PASSED=============================================

**Please note that --testfile flag defaults to test.tf and --testcondition flag defaults to testcondition.yaml and that the subcommand
expects to find both files in the current working directory. As shown above, you can override the flags and provide relative paths to the file on your local
filesystem.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		terraformruns.InitAndApplyForTesting(cwd, testfile)
		testterraform.NowProcessAndTest(cwd, testcondition)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().StringVarP(&testfile, "testfile", "", "test.tf", "Option that lets you test a *.tf file present in cwd")
	testCmd.Flags().StringVarP(&testcondition, "testcondition", "", "testcondition.yaml", "Option that takes in a yaml file with conditions to test against")

}
