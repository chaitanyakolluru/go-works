package cmd

import (
	"fmt"
	"testing"
)

func TestExtractHieraHierarchy(t *testing.T) {
	bufCH := make(chan []string, 1)
	hierarchyWant := []string{"node/%{::clientcert}", "application/%{::aa_datacenter}/%{::aa_sdlc_environment}/%{::aa_application}/%{::aa_application_tier}", "application/%{::aa_datacenter}/%{::aa_sdlc_environment}/%{::aa_application}", "application/%{::aa_sdlc_environment}/%{::aa_application}/%{::aa_application_tier}", "application/%{::aa_sdlc_environment}/%{::aa_application}", "application/common/%{::aa_datacenter}/%{::aa_application}/%{::aa_application_tier}", "application/common/%{::aa_datacenter}/%{::aa_application}", "application/common/%{::aa_application}/%{::aa_application_tier}", "application/common/%{::aa_application}", "datacenter/%{::aa_datacenter}/%{::aa_sdlc_environment}", "sdlc_environment/%{::aa_sdlc_environment}", "datacenter/%{::aa_datacenter}", "msp/%{::aa_msp}", "defaults"}
	extractHieraHierarchy(bufCH, "~/Documents/hiera.yaml")
	hierarchyGot := <-bufCH
	for i, item := range hierarchyWant {
		if item != hierarchyGot[i] {
			t.Error("ExtractHieraHierarchy function output is incorrect, got: ", hierarchyWant, " want: ", hierarchyGot)
			break
		}
	}
}

func TestGrepSearchString(t *testing.T) {
	bufK := make(chan map[string]string, 1)
	hieradata := "~/Documents/lab-control-repo/hieradata"
	searchString := "system_env"

	grepSearchString(bufK, hieradata, searchString)

	fmt.Println(<-bufK)
}
