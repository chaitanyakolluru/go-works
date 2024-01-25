package populateyaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"yamlit/cmd/clonerepo"

	"gopkg.in/yaml.v2"
)

// PopulateYaml function reads from services-offered.yaml file and creates variables-in-modues.yaml with all variables exported within each module.
func PopulateYaml(cwd, pattern, vertical, appname, componentname, storageaccount, container, key, accesskey, username, token string) {

	if pattern == "" {
		servicesOffered, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + "services-offered.yaml")
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

		clonerepo.CloneRepo(cwd, username, token)

		for key := range mapServicesOffered {
			str := fmt.Sprint(key)
			if str != "Include Service Principal Authentication in terraform code" {
				varFile, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform" + string(os.PathSeparator) + "azure-modules" + string(os.PathSeparator) + str + string(os.PathSeparator) + "variables.tf")

				temp := map[interface{}]string{}
				temp["version"] = "FILL OUT"
				//	var lineStore string
				for _, varData := range strings.Split(string(varFile), "\n") {
					ifC, _ := regexp.MatchString("variable", varData)
					if ifC {

						ifDoubleQuotes, _ := regexp.MatchString("\"", varData)
						if ifDoubleQuotes {
							inner := strings.Split(strings.Split(varData, " \"")[1], "\" ")[0]
							temp[inner] = "FILL OUT"
						} else {
							inner := strings.Split(varData, " ")[1]
							temp[inner] = "FILL OUT"
						}

					}

				}
				mapServicesOffered[str] = temp
			}

		}

		if _, ok := mapServicesOffered["Include Service Principal Authentication in terraform code"]; ok {
			mapServicesOffered["sp-auth-within-azurerm-block"] = make(map[string]string)
			mapServicesOffered["sp-auth-within-azurerm-block"].(map[string]string)["subscription_id"] = "FILL OUT"
			mapServicesOffered["sp-auth-within-azurerm-block"].(map[string]string)["tenant_id"] = "FILL OUT"
			mapServicesOffered["sp-auth-within-azurerm-block"].(map[string]string)["client_id"] = "FILL OUT"

			delete(mapServicesOffered, "Include Service Principal Authentication in terraform code")
		}

		mapServicesOffered = includeStateDetails(mapServicesOffered, pattern)

		d, _ := yaml.Marshal(&mapServicesOffered)

		dFinal := string(d) + "\n#Note: client_secret for sp based authentication to be provided as an environment variable ARM_CLIENT_SECRET."
		if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"variables-in-modules.yaml", []byte(dFinal), 0644); err != nil {
			log.Fatal(err)
		}

		if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
			log.Fatal(err)
		}

	} else {
		clonerepo.CloneRepo(cwd, username, token)

		items, err := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform" + string(os.PathSeparator) + "azure-patterns" + string(os.PathSeparator) + pattern + ".yaml")
		if err != nil {
			log.Fatal(err)
		}

		mapPatterns := make([]string, 0)
		errYml := yaml.Unmarshal([]byte(items), &mapPatterns)
		if errYml != nil {
			log.Fatal(errYml)
		}

		appInfra := make(map[interface{}]interface{})
		for _, item := range mapPatterns {
			itemS := strings.Split(item, "/")
			patFile, errPat := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform" + string(os.PathSeparator) + "azure-modules" + string(os.PathSeparator) + itemS[0] + string(os.PathSeparator) + "patterns" + string(os.PathSeparator) + itemS[1] + ".yaml")
			if errPat != nil {
				log.Fatal(errPat)
			}
			temp := make(map[string]interface{})
			errP := yaml.Unmarshal([]byte(patFile), &temp)
			if errP != nil {
				log.Fatal(errP)
			}
			for kT, vT := range temp {
				appInfra[kT] = vT
			}

		}

		for kA, vA := range appInfra {
			for kAI, vAI := range vA.(map[interface{}]interface{}) {
				if vAI == nil {
					kAIS := fmt.Sprint(kAI)
					resI := strings.Split(kAIS, "_name")[0]
					if resI == "resource_group" {
						appInfra[kA].(map[interface{}]interface{})[kAIS] = createResGrpNameFromMetadata(vertical, appname, componentname)
					} else {
						appInfra[kA].(map[interface{}]interface{})[kAIS] = createResNameFromMetadata(resI, appname, componentname)
					}

				}
			}

		}

		appInfra = includeStateDetailsPattern(appInfra, pattern, storageaccount, container, key, accesskey)

		appInfraOut, errAppI := yaml.Marshal(&appInfra)
		if errAppI != nil {
			log.Fatal(errAppI)
		}
		if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"infrastructure.yaml", []byte(appInfraOut), 0644); err != nil {
			log.Fatal(err)
		}

		if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
			log.Fatal(err)
		}

	}

}

// createResNameFromMetadata function creates names for all resources except for resource group;
// resource names have limitations around using underscores and caps; thus concatening metadata with resource names.
func createResNameFromMetadata(res, appname, componentname string) string {
	tempRes := strings.Split(res, "_")
	tempResName := ""
	for _, i := range tempRes {
		tempResName += string(i[0])
	}

	out := tempResName + appname[0:3] + componentname[0:3] + "ENVIRONMENT" + "LOCATION"
	return out
}

// createResGrpNameFromMetadata function creates names for resource group
func createResGrpNameFromMetadata(vertical, appname, componentname string) string {
	out := "aa-" + vertical + "-" + appname + "-" + componentname + "-" + "ENVIRONMENT" + "-" + "LOCATION" + "-" + "rg"
	return out
}

// includeStateDetails function adds in remote terraform state file management details to the yaml output of populate.
func includeStateDetails(mapY map[interface{}]interface{}, pattern string) map[interface{}]interface{} {
	mapY["remote-state"] = make(map[interface{}]interface{})

	mapY["remote-state"].(map[interface{}]interface{})["storage_account_name"] = "FILL OUT"
	mapY["remote-state"].(map[interface{}]interface{})["container_name"] = "FILL OUT"
	mapY["remote-state"].(map[interface{}]interface{})["key"] = "FILL OUT"
	mapY["remote-state"].(map[interface{}]interface{})["access_key"] = "FILL OUT"

	return mapY

}

// includeStateDetails function adds in remote terraform state file management details to the yaml output of populate.
func includeStateDetailsPattern(mapY map[interface{}]interface{}, pattern, storageaccount, container, key, accesskey string) map[interface{}]interface{} {
	mapY["remote-state"] = make(map[interface{}]interface{})

	// these will be obtainer from yamlit apply from of app-prerequisites.yaml file that contains instructions to create
	// appSP, appRG, add appSP as a contributor to appRG, create storage account and container for remote state mgmt of infrastructure.yaml
	mapY["remote-state"].(map[interface{}]interface{})["storage_account_name"] = storageaccount
	mapY["remote-state"].(map[interface{}]interface{})["container_name"] = container
	mapY["remote-state"].(map[interface{}]interface{})["key"] = key
	if accesskey != "" {
		mapY["remote-state"].(map[interface{}]interface{})["access_key"] = accesskey
	}

	return mapY

}
