package main

import (
	"os"
	"terraformapi/discoverservices"
	"terraformapi/populateyaml"
	"terraformapi/terraformruns"
)

func main() {
	cwd, _ := os.Getwd()
	// api that takes in yaml  to convert to terraform code, run terraform and deploy as expected

	// one api as discovery api that returns the services offered by the api
	discoverservices.DiscoverServices(cwd)

	// api that take in yaml with 1's on the resources and returns a yaml with all fields that needs populating.
	populateyaml.PopulateYaml(cwd)

	// one api that takes in that yaml and then runs init plan and returns an identifier for the call and also a plan output
	terraformruns.InitAndPlan(cwd)
	terraformruns.InitAndApply(cwd)

	// azure sdk for authentication and terraform command execution

}
