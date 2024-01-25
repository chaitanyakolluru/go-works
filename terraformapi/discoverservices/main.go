package discoverservices

import (
	"io/ioutil"
	"log"
	"os"
	"terraformapi/clonerepo"
)

func DiscoverServices(cwd string) {
	clonerepo.CloneRepo(cwd)
	items, err := ioutil.ReadDir(cwd + "\\.temp\\terraform\\azure-modules")
	if err != nil {
		log.Fatal(err)
	}

	modules := make([]string, 0)
	for _, f := range items {
		if f.IsDir() {
			modules = append(modules, f.Name())
		}
	}

	yamlWriter, err := os.OpenFile(cwd+"\\services-offered.yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer yamlWriter.Close()

	yamlLines := "Please input 1 on the resources you need creating.." + "\n" + "\n"
	for _, mod := range modules {
		yamlLines = yamlLines + mod + ": " + "\n"
	}
	yamlWriter.Write([]byte(yamlLines))

	if err := os.RemoveAll(cwd + "\\.temp"); err != nil {
		log.Fatal(err)
	}

}
