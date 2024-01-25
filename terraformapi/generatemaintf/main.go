package generatemaintf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"terraformapi/clonerepo"

	"gopkg.in/yaml.v2"
)

func GenerateMaintf(cwd string) {

	// all data in variables-in-modules.yaml needs to be on a single line for yaml decoding purposes
	variablesInModules, _ := ioutil.ReadFile(cwd + "\\variables-in-modules.yaml")
	mapVariablesInModules := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(variablesInModules), &mapVariablesInModules); err != nil {
		log.Fatal(err)
	}

	clonerepo.CloneRepo(cwd)

	mainFileWriter, err := os.OpenFile(cwd+"\\.temp\\main.tf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer mainFileWriter.Close()

	mainFileLines := ""

	mainFileLines += "provider azurerm {" + "\n"
	mainFileLines += "features {}" + "\n"
	mainFileLines += "}"
	mainFileLines += "\n"

	for key, value := range mapVariablesInModules {
		keyS := fmt.Sprint(key)
		mainFileLines += "module " + "\"" + keyS + "-mod\" {" + "\n"
		versionS := fmt.Sprint(value.(map[interface{}]interface{})["version"])
		mainFileLines += "\t" + "source = \"git@ghe.aa.com:AA/terraform.git//azure-modules/" + keyS + "?ref=" + strings.Split(strings.Split(versionS, "(")[1], ")")[1] + "\"\n"
		for keyI, valueI := range value.(map[interface{}]interface{}) {
			keyIS := fmt.Sprint(keyI)
			valueIS := fmt.Sprint(valueI)
			tempA := strings.Split(strings.Split(valueIS, "(")[1], ")")
			if keyIS != "version" {
				if tempA[0] == "string" {
					mainFileLines += "\t" + keyIS + " = " + "\"" + tempA[1] + "\"" + "\n"
				} else {
					mainFileLines += "\t" + keyIS + " = " + tempA[1] + "\n"
				}
			}

		}
		mainFileLines += "}" + "\n"
		mainFileLines += "\n"
	}

	mainFileWriter.Write([]byte(mainFileLines))

}
