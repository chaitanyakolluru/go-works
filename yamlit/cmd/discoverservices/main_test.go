package discoverservices

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestDiscoverServices(t *testing.T) {

	servicesOfferedWant := [22]string{"Include Service Principal Authentication in terraform code", "aks", "apim", "app-insights", "appservice", "azure_sql", "cdn-profile", "container-registry", "cosmosdb", "eventhub", "function-app", "keyvault", "log-analytics", "mysql_server", "resource-group", "route-table", "snet", "storage_account", "streamanalytics", "terraform-examples", "traffic-manager", "vnet"}
	servicesOfferedWantLine := ""
	for _, ll := range servicesOfferedWant {
		servicesOfferedWantLine += ll + ","
	}

	cwd, _ := os.Getwd()
	DiscoverServices(cwd, "", "")
	servicesOfferedGot := make([]string, 30)
	servicesOfferedGotBytes, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + "services-offered.yaml")
	servicesOfferedGotLine := ""
	for _, line := range strings.Split(string(servicesOfferedGotBytes), "\n") {
		ll, _ := regexp.MatchString(":", line)
		if ll {
			servicesOfferedGot = append(servicesOfferedGot, strings.Split(line, ":")[0])
			servicesOfferedGotLine += strings.Split(line, ":")[0] + ","
		}
	}

	for _, ll := range servicesOfferedWant {
		isIt := false
		for _, jj := range servicesOfferedGot {
			if ll == jj {
				isIt = true
			}
		}
		if isIt == false {
			t.Errorf("Discover command output is incorrect, got: %s, want: %s", servicesOfferedGotLine, servicesOfferedWantLine)
		}
	}

	err := os.Remove(cwd + string(os.PathSeparator) + "services-offered.yaml")
	if err != nil {
		t.Fatal(err)
	}

}
