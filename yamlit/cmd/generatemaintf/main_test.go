package generatemaintf

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateMaintf(t *testing.T) {
	cwd, _ := os.Getwd()

	// test with mocked infraastrucure.yaml file generated as the output of yamlit populate with a --pattern input.
	infraYamlContent := `
remote-state:
  container_name: remotestatecontainer
  key: terraform.tfstate
  storage_account_name: saremotestate
storage_account:
  access_tier: Hot
  account_kind: StorageV2
  account_replication_type: grs
  account_tier: Standard
  enable_https_traffic_only: true
  enable_static_website: true
  error_404_document: error.html
  index_document: index.html
  resource_group_name: aa-aot-mvp-staticwebsite-ENVIRONMENT-LOCATION-rg
  storage_account_name: samvpstaENVIRONMENTLOCATION
  version: storage-account-v1.1.2
  `
	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"infrastructure.yaml", []byte(infraYamlContent), 0644); err != nil {
		t.Error(err)
	}

	maintfWant := map[string]string{

		"container_name":            "remotestatecontainer",
		"key":                       "terraform.tfstate",
		"source":                    "storage-account-v1.1.2",
		"account_tier":              "Standard",
		"enable_static_website":     "true",
		"resource_group_name":       "aa-aot-mvp-staticwebsite-dev-eastus-rg",
		"storage_account_name":      "samvpstadeveastus",
		"access_tier":               "Hot",
		"account_kind":              "StorageV2",
		"account_replication_type":  "grs",
		"enable_https_traffic_only": "true",
		"error_404_document":        "error.html",
		"index_document":            "index.html",
	}

	GenerateMaintf(cwd, "", "dev", "eastus", "", "")
	maintfGot, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	maintfGotStr := strings.Split(string(maintfGot), "\n")
	for k, v := range maintfWant {
		present := 0
		for _, vG := range maintfGotStr {
			if lk, _ := regexp.MatchString(k, vG); lk {
				if lv, _ := regexp.MatchString(v, vG); lv {
					present = 1
				}
			}

		}
		if present == 0 {
			t.Errorf("Generate command output is incorrect,  k: %s, v: %s", k, v)
		}
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "infrastructure.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

	//test with the assumption that --pattern feature within yamlit populate is not used and generate
	// is triggered with a variables-in-modules.yaml file in the current working directory
	// and with Service Principal based authentication enabled within the yaml file.
	varModYamlContent := `
client_id: CLIENT_ID
client_secret: CLIENT_SECRET
remote-state:
  container_name: remotestatecontainer
  key: terraform.tfstate
  storage_account_name: saremotestate
storage_account:
  access_tier: Hot
  account_kind: StorageV2
  account_replication_type: grs
  account_tier: Standard
  enable_https_traffic_only: true
  enable_static_website: true
  error_404_document: error.html
  index_document: index.html
  resource_group_name: aa-aot-mvp-staticwebsite-dev-eastus-rg
  storage_account_name: samvpstadeveastus
  version: storage-account-v1.1.2
subscription_id: SUBSCRIPTION_ID
tenant_id: TENANT_ID
`
	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"variables-in-modules.yaml", []byte(varModYamlContent), 0644); err != nil {
		t.Error(err)
	}

	varModYamlWant := map[string]string{

		"container_name":            "remotestatecontainer",
		"key":                       "terraform.tfstate",
		"source":                    "storage-account-v1.1.2",
		"account_tier":              "Standard",
		"enable_static_website":     "true",
		"resource_group_name":       "aa-aot-mvp-staticwebsite-dev-eastus-rg",
		"storage_account_name":      "samvpstadeveastus",
		"access_tier":               "Hot",
		"account_kind":              "StorageV2",
		"account_replication_type":  "grs",
		"enable_https_traffic_only": "true",
		"error_404_document":        "error.html",
		"index_document":            "index.html",
		"client_id":                 "CLIENT_ID",
		"CLIENT_SECRET":             "CLIENT_SECRET",
		"subscription_id":           "SUBSCRIPTION_ID",
		"tenant_id":                 "TENANT_ID",
	}

	GenerateMaintf(cwd, "", "", "", "", "")
	varModYamlGot, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	varModYamlGotStr := strings.Split(string(varModYamlGot), "\n")
	for k, v := range varModYamlWant {
		present := 0
		for _, vG := range varModYamlGotStr {
			if lk, _ := regexp.MatchString(k, vG); lk {
				if lv, _ := regexp.MatchString(v, vG); lv {
					present = 1
				}
			}

		}
		if present == 0 {
			t.Errorf("Generate command output is incorrect,  k: %s, v: %s", k, v)
		}
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "variables-in-modules.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

	// test with the assumption that --pattern feature within yamlit populate is not used and generate
	// is triggered with a variables-in-modules.yaml file in the current working directory
	// and with Service Principal based authentication disabled within the yaml file.
	varModYamlContentNoSP := `
remote-state:
  container_name: remotestatecontainer
  key: terraform.tfstate
  storage_account_name: saremotestate
storage_account:
  access_tier: Hot
  account_kind: StorageV2
  account_replication_type: grs
  account_tier: Standard
  enable_https_traffic_only: true
  enable_static_website: true
  error_404_document: error.html
  index_document: index.html
  resource_group_name: aa-aot-mvp-staticwebsite-dev-eastus-rg
  storage_account_name: samvpstadeveastus
  version: storage-account-v1.1.2
`
	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"variables-in-modules.yaml", []byte(varModYamlContentNoSP), 0644); err != nil {
		t.Error(err)
	}

	varModYamlWantNoSP := map[string]string{

		"container_name":            "remotestatecontainer",
		"key":                       "terraform.tfstate",
		"source":                    "storage-account-v1.1.2",
		"account_tier":              "Standard",
		"enable_static_website":     "true",
		"resource_group_name":       "aa-aot-mvp-staticwebsite-dev-eastus-rg",
		"storage_account_name":      "samvpstadeveastus",
		"access_tier":               "Hot",
		"account_kind":              "StorageV2",
		"account_replication_type":  "grs",
		"enable_https_traffic_only": "true",
		"error_404_document":        "error.html",
		"index_document":            "index.html",
	}

	GenerateMaintf(cwd, "", "", "", "", "")
	varModYamlGotNoSP, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	varModYamlGotStrNoSP := strings.Split(string(varModYamlGotNoSP), "\n")
	for k, v := range varModYamlWantNoSP {
		present := 0
		for _, vG := range varModYamlGotStrNoSP {
			if lk, _ := regexp.MatchString(k, vG); lk {
				if lv, _ := regexp.MatchString(v, vG); lv {
					present = 1
				}
			}

		}
		if present == 0 {
			t.Errorf("Generate command output is incorrect,  k: %s, v: %s", k, v)
		}
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "variables-in-modules.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

	// test if on specifying a --inputyaml flag to generate creates terraform code based upon the input yaml file
	varModYamlContentInputYaml := `
remote-state:
  container_name: remotestatecontainer
  key: terraform.tfstate
  storage_account_name: saremotestate
storage_account:
  access_tier: Hot
  account_kind: StorageV2
  account_replication_type: grs
  account_tier: Standard
  enable_https_traffic_only: true
  enable_static_website: true
  error_404_document: error.html
  index_document: index.html
  resource_group_name: aa-aot-mvp-staticwebsite-dev-eastus-rg
  storage_account_name: samvpstadeveastus
  version: storage-account-v1.1.2
`
	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"testYml.yaml", []byte(varModYamlContentInputYaml), 0644); err != nil {
		t.Error(err)
	}

	varModYamlWantInputYaml := map[string]string{

		"container_name":            "remotestatecontainer",
		"key":                       "terraform.tfstate",
		"source":                    "storage-account-v1.1.2",
		"account_tier":              "Standard",
		"enable_static_website":     "true",
		"resource_group_name":       "aa-aot-mvp-staticwebsite-dev-eastus-rg",
		"storage_account_name":      "samvpstadeveastus",
		"access_tier":               "Hot",
		"account_kind":              "StorageV2",
		"account_replication_type":  "grs",
		"enable_https_traffic_only": "true",
		"error_404_document":        "error.html",
		"index_document":            "index.html",
	}

	GenerateMaintf(cwd, "testYml.yaml", "", "", "", "")
	varModYamlGotInputYaml, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	varModYamlGotStrInputYaml := strings.Split(string(varModYamlGotInputYaml), "\n")
	for k, v := range varModYamlWantInputYaml {
		present := 0
		for _, vG := range varModYamlGotStrInputYaml {
			if lk, _ := regexp.MatchString(k, vG); lk {
				if lv, _ := regexp.MatchString(v, vG); lv {
					present = 1
				}
			}

		}
		if present == 0 {
			t.Errorf("Generate command output is incorrect,  k: %s, v: %s", k, v)
		}
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "testYml.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

	// test if input file contains maps as values, with arrays within them containing maps and lists with maps within them (using incorrect yaml file for testing for all scenarios)
	varModYamlComplex1 := `
remote-state:
    container_name: remotestatecontainer
    key: terraform.tfstate
    storage_account_name: saremotestate
appservice:
    appservice_plan_name: demo-asp
    appservice_plan_tier: Standard
    appservice_plan_size: S1
    appservice_plan_kind: Linux
    new_dummy:
    - x: dummy1
      a: dummy2
    resource_group_name: demo-rg
    appservice_plan_reserved: true
    appservice:
      appservice-1:
        appservice_name: demo-appservice
        site_config:
        - always_on: true
          app_command_line: "some command"
          linux_fx_version: JAVA|11-java11
    version: appservice-v2.3.1
`
	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"infrastructure.yaml", []byte(varModYamlComplex1), 0644); err != nil {
		t.Error(err)
	}

	varModYamlComplex1Want := map[string]string{

		"container_name":           "remotestatecontainer",
		"key":                      "terraform.tfstate",
		"storage_account_name":     "saremotestate",
		"appservice_plan_name":     "demo-asp",
		"appservice_plan_tier":     "Standard",
		"appservice_plan_size":     "S1",
		"appservice_plan_kind":     "Linux",
		"resource_group_name":      "demo-rg",
		"appservice_plan_reserved": "true",
		"appservice_name":          "demo-appservice",
		"always_on":                "true",
		"app_command_line":         "some command",
		"linux_fx_version":         "JAVA|11-java11",
		"source":                   "appservice-v2.3.1",
		"x":                        "dummy1",
		"a":                        "dummy2",
	}

	GenerateMaintf(cwd, "", "dev", "eastus", "", "")
	varModYamlComplex1Got, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "main.tf")
	varModYamlComplex1GotStr := strings.Split(string(varModYamlComplex1Got), "\n")
	for k, v := range varModYamlComplex1Want {
		present := 0
		for _, vG := range varModYamlComplex1GotStr {
			if lk, _ := regexp.MatchString(k, vG); lk {
				if lv, _ := regexp.MatchString(v, vG); lv {
					present = 1
				}
			}

		}
		if present == 0 {
			t.Errorf("Generate command output is incorrect,  k: %s, v: %s", k, v)
		}
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "infrastructure.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}
