package terraformruns

import (
	"log"
	"os"
	"terraformapi/generatemaintf"
	"terraformapi/terraformfunctions"
)

func InitAndApply(cwd string) {
	generatemaintf.GenerateMaintf(cwd)
	if err := os.RemoveAll(cwd + "\\.temp\\terraform"); err != nil {
		log.Fatal(err)
	}

	// running init
	terraformfunctions.TerraformRuns("init", cwd)

	//apply run
	terraformfunctions.TerraformRunsApply(cwd)
}

func InitAndPlan(cwd string) {
	generatemaintf.GenerateMaintf(cwd)
	if err := os.RemoveAll(cwd + "\\.temp\\terraform"); err != nil {
		log.Fatal(err)
	}

	// running init
	terraformfunctions.TerraformRuns("init", cwd)

	//running plan
	terraformfunctions.TerraformRuns("plan", cwd)

}
