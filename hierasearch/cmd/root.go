/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// go run main.go  --searchString profiles::java --hieraFile ~/Documents/hiera.yaml --hieradata ~/Documents/lab-control-repo/hieradata

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var searchString, hieraFile, hieradata string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hierasearch",
	Short: "Searches heira and shows op based on hiera hierarchy",
	Long: `hierasearch is used to take in a search string and give
	results based on hiera hierarchy`,

	Run: func(cmd *cobra.Command, args []string) {

		shouldExit := 0
		if searchString == "" {
			fmt.Println("--searchString is not defined")
			shouldExit = 1
		}
		if _, err := os.Stat(hieraFile); os.IsNotExist(err) {
			fmt.Printf("--hieraFile %s path doesn't exist", hieraFile)
			shouldExit = 1
		}
		if _, err := os.Stat(hieradata); os.IsNotExist(err) {
			fmt.Printf("--hieradata %s path doesn't exist\n", hieradata)
		}

		if shouldExit == 1 {
			cmd.Usage()
			os.Exit(1)
		}
		bufCH := make(chan []string, 1)
		bufK := make(chan map[string]string, 1)
		bufA := make(chan map[string]interface{}, 1)
		bufR := make(chan string, 100)

		go extractHieraHierarchy(bufCH, hieraFile)
		go grepSearchString(bufK, hieradata, searchString)
		go parseYaml(bufK, bufA, searchString)

		go nowImplementHieraLookup(bufCH, bufA, bufR)
		printResultsAsTheyAreAvailable(bufR)
		close(bufCH)
		close(bufK)
		close(bufA)
		close(bufR)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&searchString, "searchString", "", "", "hiera key to search with")
	rootCmd.PersistentFlags().StringVarP(&hieraFile, "hieraFile", "", "/etc/puppetlabs/hiera/hiera.yaml", "hiera.yaml file location")
	rootCmd.PersistentFlags().StringVarP(&hieradata, "hieradata", "", "."+string(os.PathSeparator)+"hieradata", "hieradata directory path")

}

func extractHieraHierarchy(bufCH chan []string, hieraFile string) {
	// hieraFileText, _ := ioutil.ReadFile(hieraFile)
	hieraFileText := `# managed by puppet
:hierarchy:
  - "node/%{::clientcert}"
  - "application/%{::aa_datacenter}/%{::aa_sdlc_environment}/%{::aa_application}/%{::aa_application_tier}"
  - "application/%{::aa_datacenter}/%{::aa_sdlc_environment}/%{::aa_application}"
  - "application/%{::aa_sdlc_environment}/%{::aa_application}/%{::aa_application_tier}"
  - "application/%{::aa_sdlc_environment}/%{::aa_application}"
  - "application/common/%{::aa_datacenter}/%{::aa_application}/%{::aa_application_tier}"
  - "application/common/%{::aa_datacenter}/%{::aa_application}"
  - "application/common/%{::aa_application}/%{::aa_application_tier}"
  - "application/common/%{::aa_application}"
  - "datacenter/%{::aa_datacenter}/%{::aa_sdlc_environment}"
  - "sdlc_environment/%{::aa_sdlc_environment}"
  - "datacenter/%{::aa_datacenter}"
  - "msp/%{::aa_msp}"
  - defaults
`

	hierarchy := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(hieraFileText), &hierarchy)
	if err != nil {
		log.Fatal(err)
	}

	hiera := make([]string, 0)
	for _, item := range hierarchy[":hierarchy"].([]interface{}) {
		hiera = append(hiera, item.(string))
	}

	bufCH <- hiera

}

func grepSearchString(bufK chan map[string]string, hieradata string, searchString string) {
	grepCmd := exec.Command("grep", "-r", searchString, hieradata)
	grepCmdOut, err := grepCmd.Output()
	if err != nil {
		fmt.Println("No results found!!", err, grepCmdOut)
	}
	yamlFiles := map[string]string{}
	for _, line := range strings.Split(string(grepCmdOut), "\n") {
		if line != "" {
			yamlFiles[strings.Split(line, ":")[0]] = ""
		}
	}
	bufK <- yamlFiles

}

func parseYaml(bufK chan map[string]string, bufA chan map[string]interface{}, searchString string) {
	yamlFiles := <-bufK
	actualYaml := make(map[string]interface{})
	for yFile := range yamlFiles {
		actualYaml[yFile] = make([]map[string]interface{}, 0)
		yFileContent, _ := ioutil.ReadFile(yFile)
		yFileContentYaml := make(map[string]interface{})
		if err := yaml.Unmarshal([]byte(yFileContent), &yFileContentYaml); err != nil {
			log.Fatal(err)
		}
		for line, value := range yFileContentYaml {
			r, _ := regexp.Compile("^#.*")

			if !r.MatchString(line) {
				ifC, _ := regexp.MatchString(searchString, line)
				if ifC {
					lineValue := make(map[string]interface{})
					lineValue[line] = value
					actualYaml[yFile] = append(actualYaml[yFile].([]map[string]interface{}), lineValue)
				}
			}

		}
		if len(actualYaml[yFile].([]map[string]interface{})) == 0 {
			delete(actualYaml, yFile)
		}
	}
	bufA <- actualYaml

}

func nowImplementHieraLookup(bufCH chan []string, bufA chan map[string]interface{}, bufR chan string) {
	hierarchy := <-bufCH
	actualYaml := <-bufA

	datacenter := []string{"lab", "pdc", "cdc", "lab_evs", "dfw", "phx", "dxc_pdc", "dxc_cdc"}
	sdlcEnvironment := []string{"development", "test", "stage", "production"}
	msp := []string{"aot", "evs", "dxc", "ibm"}

	for _, hie := range hierarchy {
		hieSplit := strings.Split(hie, "/")
		for file, item := range actualYaml {
			file = strings.TrimRight(file, ".yaml")
			fileSplit := strings.Split(strings.Split(file, "/hieradata/")[1], "/")
			if len(hieSplit) == len(fileSplit) {
				opLine := ""
				for i, hieS := range hieSplit {
					match := 0
					r, _ := regexp.Compile("%{::")
					if r.MatchString(hieS) {

						switch hieS {
						case "%{::aa_sdlc_environment}":
							for _, sd := range sdlcEnvironment {
								if sd == fileSplit[i] {
									match = 1
									opLine += fmt.Sprintf("For SLDC: %s, ", strings.ToUpper(sd))
								}
							}

						case "%{::aa_datacenter}":
							for _, dc := range datacenter {
								if dc == fileSplit[i] {
									match = 1
									opLine += fmt.Sprintf("For DC: %s, ", strings.ToUpper(dc))
								}
							}

						case "%{::aa_msp}":
							for _, ms := range msp {
								if ms == fileSplit[i] {
									match = 1
									opLine += fmt.Sprintf("For MSP: %s, ", strings.ToUpper(ms))
								}
							}

						default:
							for i = len(hieSplit) - 1; i >= 0; i-- {
								switch hieSplit[i] {
								case "%{::clientcert}":
									opLine += fmt.Sprintf("For Certname: %s, ", strings.ToUpper(fileSplit[i]))
								case "%{::aa_application_tier}":
									opLine += fmt.Sprintf("For App_Tier: %s, ", strings.ToUpper(fileSplit[i]))
								case "%{::aa_application}":
									opLine += fmt.Sprintf("For App: %s, ", strings.ToUpper(fileSplit[i]))
								default:
									continue
								}
							}
							itemYaml, errYaml := yaml.Marshal(&item)
							if errYaml != nil {
								log.Fatal(errYaml)
							}
							bufR <- fmt.Sprint(strings.TrimRight(opLine, ", "), "\n\n", string(itemYaml))
						}

					} else {
						if hieS == fileSplit[i] {
							match = 1
							if len(hieSplit) == 1 {
								opLine += fmt.Sprintf("For: %s", strings.ToUpper(hieS))
							}
						}
					}
					if match == 1 {
						if i == len(hieSplit)-1 {
							itemYaml, errYaml := yaml.Marshal(&item)
							if errYaml != nil {
								log.Fatal(errYaml)
							}
							bufR <- fmt.Sprint(strings.TrimRight(opLine, ", "), "\n\n", string(itemYaml))
						}
					} else {
						break
					}

				}
			}

		}
	}

	// signling to printResultsAsTheyAreAvailable end of stream
	bufR <- "EOS"
}

func printResultsAsTheyAreAvailable(bufR chan string) {
	for {
		printItem, res := <-bufR
		if !res {
			return
		}
		if printItem != "EOS" {
			fmt.Println(printItem)
		} else {
			return
		}

	}
}
