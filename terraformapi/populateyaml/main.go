package populateyaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"terraformapi/clonerepo"

	"gopkg.in/yaml.v2"
)

func PopulateYaml(cwd string) {
	servicesOffered, _ := ioutil.ReadFile(cwd + "\\services-offered.yaml")
	mainString := ""
	for _, line := range strings.Split(string(servicesOffered), "\n") {
		ll, _ := regexp.MatchString(":", line)
		if ll {
			mainString = mainString + line + "\n"
		}
	}

	mapServicesOffered := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(mainString), &mapServicesOffered)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range mapServicesOffered {
		if value == nil {
			delete(mapServicesOffered, key)

		}
	}

	clonerepo.CloneRepo(cwd)

	for key := range mapServicesOffered {
		str := fmt.Sprint(key)
		varFile, _ := ioutil.ReadFile(cwd + "\\.temp\\terraform\\azure-modules\\" + str + "\\variables.tf")

		temp := map[interface{}]string{}
		temp["version"] = "(string)"
		var lineStore string
		for _, varData := range strings.Split(string(varFile), "\n") {
			ifC, _ := regexp.MatchString("variable", varData)
			if ifC {
				inner := strings.Split(strings.Split(varData, " \"")[1], "\" ")[0]
				temp[inner] = ""
				lineStore = inner
			}

			ifCi, _ := regexp.MatchString("type", varData)
			ifCi2, _ := regexp.MatchString("description", varData)
			ifCi3, _ := regexp.MatchString("variable", varData)

			if ifCi && !ifCi2 && !ifCi3 {
				inner2 := strings.TrimSpace(strings.Split(varData, "=")[1])
				temp[lineStore] = "(" + inner2 + ")"
			}

		}
		mapServicesOffered[str] = temp
	}

	d, _ := yaml.Marshal(&mapServicesOffered)
	if err := ioutil.WriteFile(cwd+"\\variables-in-modules.yaml", []byte(d), 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(cwd + "\\.temp"); err != nil {
		log.Fatal(err)
	}

}
