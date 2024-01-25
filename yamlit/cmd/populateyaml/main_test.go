package populateyaml

import (
	"bytes"
	"io/ioutil"
	"os"

	"strings"
	"testing"
)

func TestPopulateYaml(t *testing.T) {
	cwd, _ := os.Getwd()


	// test regular usage of populate command.
	servicesOfferedYaml := `
Please select the items needed by placing a 1 beside the item..

Include Service Principal Authentication in terraform code: 1
	
aks: 
apim: 
app-insights: 
appservice: 1
appservice-cert-hostname-binding: 1
azure_sql: 
cdn-profile: 
container-registry: 
cosmosdb: 
eventhub: 
function-app: 
keyvault: 
log-analytics: 
monitor-diagnostics-settings: 
mysql_server: 
resource-group: 
route-table: 
snet: 
storage_account:
streamanalytics: 
terraform-examples: 
traffic-manager: 
vnet: 	
	`

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"services-offered.yaml", []byte(servicesOfferedYaml), 0644); err != nil {
		t.Fatal(err)
	}

	outWantStr := `appservice:
  appservice: FILL OUT
  appservice_plan_kind: FILL OUT
  appservice_plan_name: FILL OUT
  appservice_plan_reserved: FILL OUT
  appservice_plan_size: FILL OUT
  appservice_plan_tier: FILL OUT
  appservice_slot: FILL OUT
  custom_tags: FILL OUT
  dynamic_app_settings: FILL OUT
  existing_appservice_plan_name: FILL OUT
  identity_type: FILL OUT
  keyvault_name: FILL OUT
  resource_group_name: FILL OUT
  version: FILL OUT
appservice-cert-hostname-binding:
  app_service_name: FILL OUT
  app_service_plan_name: FILL OUT
  hostname: FILL OUT
  resource_group_name: FILL OUT
  version: FILL OUT
remote-state:
  access_key: FILL OUT
  container_name: FILL OUT
  key: FILL OUT
  storage_account_name: FILL OUT
sp-auth-within-azurerm-block:
  client_id: FILL OUT
  subscription_id: FILL OUT
  tenant_id: FILL OUT

#Note: client_secret for sp based authentication to be provided as an environment variable ARM_CLIENT_SECRET.`
	outWant := []byte(outWantStr)

	PopulateYaml(cwd, "", "", "", "", "", "", "", "", "", "")

	outGot, errOG := ioutil.ReadFile(cwd + string(os.PathSeparator) + "variables-in-modules.yaml")
	if errOG != nil {
		t.Fatal(errOG)
	}

	if !bytes.Equal(outWant, outGot) {
		t.Errorf("Populate command output is incorrect, got: %s, want: %s", string(outGot), outWantStr)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "services-offered.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "variables-in-modules.yaml"); err != nil {
		t.Fatal(err)
	}

	// test --pattern implementation of populate command.
	outWantStrPattern := `remote-state:
  container_name: testc
  key: testkey
  storage_account_name: testsa
storage_account:
  access_tier: Hot
  account_kind: StorageV2
  account_replication_type: grs
  account_tier: Standard
  enable_https_traffic_only: true
  enable_static_website: true
  error_404_document: error.html
  index_document: index.html
  network_rules_bypass:
  - AzureServices
  resource_group_name: aa-aot-yamlit-populate-test-ENVIRONMENT-LOCATION-rg
  storage_account_name: sayampopENVIRONMENTLOCATION
  version: storage-account-v1.1.2`
	outWantPattern := []byte(outWantStrPattern)

	PopulateYaml(cwd, "simple-reactjs-app", "aot", "yamlit", "populate-test", "testsa", "testc", "testkey", "", "", "")

	outGotPattern, errOGPattern := ioutil.ReadFile(cwd + string(os.PathSeparator) + "infrastructure.yaml")
	if errOGPattern != nil {
		t.Fatal(errOGPattern)
	}

	outGotPatternTrim := strings.TrimSpace(string(outGotPattern))
	if !bytes.Equal(outWantPattern, []byte(outGotPatternTrim)) {
		t.Errorf("Populate command output is incorrect, got: %s, want: %s", string(outGotPatternTrim), outWantStrPattern)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "infrastructure.yaml"); err != nil {
		t.Fatal(err)
	}

}
