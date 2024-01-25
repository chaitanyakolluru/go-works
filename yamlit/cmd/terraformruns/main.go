package terraformruns

import (
	"io"
	"log"
	"os"

	"yamlit/cmd/generatemaintf"
	"yamlit/cmd/terraformfunctions"
)

// InitAndApply function that runs terraform init and apply
func InitAndApply(cwd, inputyaml, environment, location, username, token string) {

	generatemaintf.GenerateMaintf(cwd, inputyaml, environment, location, username, token)
	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform"); err != nil {
		log.Fatal(err)
	}

	// running init
	terraformfunctions.TerraformRuns(cwd, "init")

	//apply run
	terraformfunctions.TerraformRuns(cwd, "apply", "-auto-approve")

}

// InitAndPlan function that runs terraform init and plan
func InitAndPlan(cwd, inputyaml, environment, location, username, token string) {

	generatemaintf.GenerateMaintf(cwd, inputyaml, environment, location, username, token)
	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform"); err != nil {
		log.Fatal(err)
	}

	// running init
	terraformfunctions.TerraformRuns(cwd, "init")

	//running plan
	terraformfunctions.TerraformRuns(cwd, "plan")

}

// InitAndApplyForTesting function that runs terraform init and apply after creating a temp directory space and copying the --testfile to it.
func InitAndApplyForTesting(cwd, testFile string) {

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(cwd+string(os.PathSeparator)+".temp", 0755); err != nil {
		log.Fatal(err)
	}
	ss, err := os.Open(cwd + string(os.PathSeparator) + testFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dd, err := os.Create(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(dd, ss)
	dd.Close()
	ss.Close()

	// running init
	terraformfunctions.TerraformRuns(cwd, "init")

	//apply run
	terraformfunctions.TerraformRuns(cwd, "apply", "-auto-approve")

}
