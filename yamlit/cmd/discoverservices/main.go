package discoverservices

import (
	"io/ioutil"
	"log"
	"os"

	"yamlit/cmd/clonerepo"
)

// DiscoverServices function discover all services offered by AA/Terraform.
func DiscoverServices(cwd, username, token string) {
	clonerepo.CloneRepo(cwd, username, token)
	items, err := ioutil.ReadDir(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform" + string(os.PathSeparator) + "azure-modules")
	if err != nil {
		log.Fatal(err)
	}

	modules := make([]string, 0)
	for _, f := range items {
		if f.IsDir() {
			modules = append(modules, f.Name())
		}
	}

	yamlWriter, err := os.OpenFile(cwd+string(os.PathSeparator)+"services-offered.yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer yamlWriter.Close()

	yamlLines := "Please select the items needed by placing a 1 beside the item.." + "\n" + "\n"
	yamlLines += "Include Service Principal Authentication in terraform code: " + "\n" + "\n"
	for _, mod := range modules {
		yamlLines = yamlLines + mod + ": " + "\n"
	}
	yamlWriter.Write([]byte(yamlLines))

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		log.Fatal(err)
	}

}
