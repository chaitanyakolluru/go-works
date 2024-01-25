package generatemaintf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"

	"yamlit/cmd/clonerepo"

	"gopkg.in/yaml.v2"
)

// GenerateMaintf function reads variables-in-modules.yaml and generates a main.tf populated with all variables appropriately.
func GenerateMaintf(cwd, inputYaml, environment, location, username, token string) {

	if inputYaml == "" {
		if environment != "" && location != "" {
			inputYaml = "infrastructure.yaml"
		} else {
			inputYaml = "variables-in-modules.yaml"
		}

	}
	variablesInModules, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + inputYaml)
	mapVariablesInModules := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(variablesInModules), &mapVariablesInModules); err != nil {
		log.Fatal(err)
	}

	clonerepo.CloneRepo(cwd, username, token)

	mainFileWriter, err := os.OpenFile(cwd+string(os.PathSeparator)+".temp"+string(os.PathSeparator)+"main.tf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer mainFileWriter.Close()

	mainFileLines := ""

	mainFileLines += "provider azurerm {" + "\n"

	if _, ok := mapVariablesInModules["subscription_id"]; ok {
		mainFileLines += "\t" + "subscription_id = \"" + fmt.Sprint(mapVariablesInModules["subscription_id"]) + "\"" + "\n"
		delete(mapVariablesInModules, "subscription_id")
	}

	if _, ok := mapVariablesInModules["client_id"]; ok {
		mainFileLines += "\t" + "client_id = \"" + fmt.Sprint(mapVariablesInModules["client_id"]) + "\"" + "\n"
		delete(mapVariablesInModules, "client_id")
	}

	if _, ok := mapVariablesInModules["client_secret"]; ok {
		mainFileLines += "\t" + "client_secret = \"" + fmt.Sprint(mapVariablesInModules["client_secret"]) + "\"" + "\n"
		delete(mapVariablesInModules, "client_secret")
	}

	if _, ok := mapVariablesInModules["tenant_id"]; ok {
		mainFileLines += "\t" + "tenant_id = \"" + fmt.Sprint(mapVariablesInModules["tenant_id"]) + "\"" + "\n"
		delete(mapVariablesInModules, "tenant_id")
	}

	mainFileLines += "\t" + "features {}" + "\n"
	mainFileLines += "}" + "\n" + "\n"

	username, token = clonerepo.GenerateAuthDetails(username, token)
	for key, value := range mapVariablesInModules {
		keyS := fmt.Sprint(key)
		if keyS == "remote-state" {
			tempRS := make(map[string]string)
			for kRS, vRS := range mapVariablesInModules["remote-state"].(map[interface{}]interface{}) {
				tempRS[fmt.Sprint(kRS)] = fmt.Sprint(vRS)
			}
			mainFileLines += "terraform {" + "\n"
			mainFileLines += "\t" + "backend \"azurerm\" {" + "\n"
			mainFileLines += "\t" + "\t" + "resource_group_name = \"" + tempRS["resource_group_name"] + "\"" + "\n"
			mainFileLines += "\t" + "\t" + "storage_account_name = \"" + tempRS["storage_account_name"] + "\"" + "\n"
			mainFileLines += "\t" + "\t" + "container_name = \"" + tempRS["container_name"] + "\"" + "\n"
			mainFileLines += "\t" + "\t" + "key = \"" + tempRS["key"] + "\"" + "\n"
			if _, ok := tempRS["access_key"]; ok {
				mainFileLines += "\t" + "\t" + "access_key = \"" + tempRS["access_key"] + "\"" + "\n"
			}

			mainFileLines += "\t" + "}" + "\n"
			mainFileLines += "}" + "\n" + "\n"
		} else {
			mainFileLines += "module " + "\"" + keyS + "-mod\" {" + "\n"
			versionS := fmt.Sprint(value.(map[interface{}]interface{})["version"])
			mainFileLines += "\t" + "source = \"git::https://" + username + ":" + token + "@ghe.aa.com/AA/terraform.git//azure-modules/" + keyS + "?ref=" + versionS + "\"\n"
			for keyI, valueI := range value.(map[interface{}]interface{}) {

				keyIS := fmt.Sprint(keyI)
				d := make([]byte, 0)
				if keyIS != "version" {

					switch valueII := reflect.ValueOf(valueI); valueII.Kind() {
					case reflect.Map:
						d, _ = json.MarshalIndent(convertMapJsonify(valueI), "\t", "\t")
					case reflect.Slice:
						d, _ = json.MarshalIndent(convertSliceJsonify(valueI), "\t", "\t")
					default:
						d, _ = json.MarshalIndent(valueI, "\t", "\t")
					}
					mENV := regexp.MustCompile(`ENVIRONMENT`)
					mLOC := regexp.MustCompile(`LOCATION`)
					tempD := mLOC.ReplaceAllString(mENV.ReplaceAllString(string(d), environment), location)

					mainFileLines += "\t" + keyIS + " = " + tempD + "\n"
				}
			}
			mainFileLines += "}" + "\n" + "\n"
		}

	}

	mainFileWriter.Write([]byte(mainFileLines))

}

// convertMapJsonify function converts map interface{} objects to strings so json.MarshalIndent can process them; Marshal function needs keys as strings and will NOT accept interface{}
func convertMapJsonify(vv interface{}) map[string]interface{} {
	vvI := make(map[string]interface{})
	for ki, vi := range vv.(map[interface{}]interface{}) {
		kiS := fmt.Sprint(ki)
		if reflect.ValueOf(vi).Kind() == reflect.Map {
			vvI[kiS] = convertMapJsonify(vi)
		} else if reflect.ValueOf(vi).Kind() == reflect.Slice {
			vvI[kiS] = convertSliceJsonify(vi)
		} else {
			vvI[kiS] = vi
		}
	}

	return vvI
}

// convertSliceJsonify function checks if elements within the slice are maps and invokes convertMapJsonify to jsonify the maps.
func convertSliceJsonify(vv interface{}) []interface{} {
	vvII := make([]interface{}, 0)
	for _, viItem := range vv.([]interface{}) {
		if reflect.ValueOf(viItem).Kind() == reflect.Map {
			vvII = append(vvII, convertMapJsonify(viItem))
		} else if reflect.ValueOf(viItem).Kind() == reflect.Slice {
			vvII = append(vvII, convertSliceJsonify(viItem))
		} else {
			vvII = append(vvII, viItem)
		}
	}

	return vvII
}
