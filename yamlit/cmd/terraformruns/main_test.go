package terraformruns

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"
)

func TestInitAndPlan(t *testing.T) {
	cwd, _ := os.Getwd()

	// test with plan run with --environment and --location as input arguments on infrastructure.yaml
	infraYamlContent := `
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
  resource_group_name: aa-aot-sais-nonprod-eastus-rg
  storage_account_name: demochai0622
  version: storage-account-v1.1.2
`

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"infrastructure.yaml", []byte(infraYamlContent), 0644); err != nil {
		t.Error(err)
	}

	outGot := captureStdout(InitAndPlan)
	outWant := "1 to add, 0 to change, 0 to destroy"
	if m, _ := regexp.MatchString(outWant, outGot); !m {
		t.Errorf("Plan output is incorrect, got: %s, want: %s", outGot, outWant)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "infrastructure.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}

func TestInitAndApply(t *testing.T) {
	cwd, _ := os.Getwd()

	// test with plan run with --environment and --location as input arguments on infrastructure.yaml
	infraYamlContent := `
remote-state:
    container_name: remotestatecontainer
    key: yamlit-testing.tfstate
    storage_account_name: 852047sa
    resource_group_name: aa-aot-sais-nonprod-eastus-rg
appservice:
    appservice_plan_name: yamlit-testing-ap
    appservice_plan_tier: Free
    appservice_plan_size: F1
    appservice_plan_kind: Linux
    resource_group_name: aa-aot-sais-nonprod-eastus-rg
    appservice_plan_reserved: true
    appservice:
      appservice-1:
        appservice_name: yamlit-testing-app-serv
    version: appservice-v2.3.1
`

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"infrastructure.yaml", []byte(infraYamlContent), 0644); err != nil {
		t.Error(err)
	}

	outGot := captureStdout(InitAndApply)
	outWant := "Apply complete!"
	if m, _ := regexp.MatchString(outWant, outGot); !m {
		t.Errorf("Plan output is incorrect, got: %s, want: %s", outGot, outWant)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "infrastructure.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}

func captureStdout(f func(string, string, string, string, string, string)) string {
	cwd, _ := os.Getwd()

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = w
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	wg.Wait()

	f(cwd, "", "dev", "eastus", "", "")

	w.Close()

	return <-out

}
