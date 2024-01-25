package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func fileExist(file string) bool {
	result := true
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			result = false
		}
	}
	return result
}

func processCacerts(cacertsListCacerts []string, aaRootCrt string, prefix string) string {
	printLine := "OK"
	aaRootCertAddedCacerts := make([]string, 0)
	keytoolPath := ""

	for _, file := range cacertsListCacerts {
		if !fileExist("/usr/bin/keytool") {
			if fileExist("/usr/java/latest/jre/bin/keytool") {
				keytoolPath = "/usr/java/latest/jre/bin/keytool"
			} else {
				keytoolPathCacerts := make([]string, 0)
				cwd, _ := os.Getwd()
				err := filepath.Walk(cwd,
					func(path string, info os.FileInfo, err error) error {
						if err != nil {
							fmt.Println("Error walking..")
							os.Exit(3)
						}
						pathCuts := strings.Split(path, "\\")
						if pathCuts[len(pathCuts)-1] == "keytool" {
							keytoolPathCacerts = append(keytoolPathCacerts, path)
						}

						return nil
					})
				if err != nil {
					fmt.Println("Error walking..")
					os.Exit(3)
				}
				keytoolPath = keytoolPathCacerts[0]
			}
		} else {
			keytoolPath = "keytool"
		}

		execCmd := exec.Command("cat", keytoolPath)
		execOut, err := execCmd.Output()
		if err != nil {
			fmt.Println("cant run cat")
			os.Exit(10)
		}

		if1, _ := regexp.MatchString("Certificate was added to keystore", string(execOut))
		if2, _ := regexp.MatchString("already exists", string(execOut))
		if if1 || if2 {
			aaRootCertAddedCacerts = append(aaRootCertAddedCacerts, file)

		}
	}

	diffSets := diffArrays(cacertsListCacerts, aaRootCertAddedCacerts)
	if len(diffSets) != 0 {
		printLine = "Not-OK: Script was not able to add the root ca " + aaRootCrt + " into these keystores: " + string(diffSets[0])

	}

	return printLine
}

func diffArrays(cacertsListCacerts []string, aaRootCertAddedCacerts []string) []string {
	type void struct{}
	var member void
	result := make([]string, 0)

	cacertsListCacertsMap := make(map[string]void)
	aaRootCertAddedCacertsMap := make(map[string]void)

	for _, item := range cacertsListCacerts {
		cacertsListCacertsMap[item] = member
	}

	for _, item := range aaRootCertAddedCacerts {
		aaRootCertAddedCacertsMap[item] = member
	}

	for item := range cacertsListCacertsMap {
		if _, ok := aaRootCertAddedCacertsMap[item]; !ok {
			result = append(result, item)
		}
	}

	return result

}

func processTrustP12(cacertsListTrustP12 []string, aaRootCrt string, prefix string) string {
	printLine := "OK"
	return printLine
}

func listAndCreateFilesZip(file *zip.File, cwd string) string {
	fileread, err := file.Open()
	if err != nil {
		fmt.Println("unable to extract zip")
		os.Exit(4)
	}

	filePath := cwd + "\\" + file.Name
	filePp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("unable to create pem file pointer for write")
		os.Exit(7)
	}
	defer filePp.Close()

	if _, err := io.Copy(filePp, fileread); err != nil {
		fmt.Println("cant write into the pem file")
		os.Exit(8)
	}

	defer fileread.Close()
	return file.Name
}

func processCaFile(cacertsList map[string][]string, caFile string) {
	listFilesZip := make([]string, 0)
	cwd, _ := os.Getwd()
	aaRootCrtZip := cwd + "\\" + caFile

	prefixForControlFile := strings.Split(strings.Split(strings.Split(caFile, ".")[0], "_")[0], "DEV")[0]

	readZip, err := zip.OpenReader(aaRootCrtZip)
	if err != nil {
		fmt.Println("cannot read from the zip")
		os.Exit(5)
	}
	defer readZip.Close()

	for _, file := range readZip.File {
		listFilesZip = append(listFilesZip, listAndCreateFilesZip(file, cwd))
	}

	for _, aaRootCrt := range listFilesZip {
		printLine := processCacerts(cacertsList["cacerts"], aaRootCrt, strings.Split(aaRootCrt, "_")[0]) + "|" + processTrustP12(cacertsList["trust.p12"], aaRootCrt, strings.Split(aaRootCrt, "_")[0])

		if printLine != "OK|OK" {
			if strings.Split(printLine, "|")[0] != "OK" {
				if strings.Split(printLine, "|")[1] == "OK" {
					printLine = strings.Split(printLine, "|")[0]
				}
			} else {
				printLine = strings.Split(printLine, "|")[1]
			}
			if err := ioutil.WriteFile(cwd+"\\"+prefixForControlFile+"_CA_NOT_ADDED", []byte(printLine), 0644); err != nil {
				fmt.Println("unable to generate not added file")
				os.Exit(7)
			}
		}

	}

	aaRootCrtZipArray := make([]string, 1)
	aaRootCrtZipArray[0] = aaRootCrtZip
	deleteFiles(aaRootCrtZipArray)
	listFilesZipPathed := make([]string, 0)
	for _, ff := range listFilesZip {
		listFilesZipPathed = append(listFilesZipPathed, cwd+"\\"+ff)
	}
	deleteFiles(listFilesZipPathed)

}

func deleteFiles(cleanupArray []string) {
	for _, file := range cleanupArray {
		os.Remove(file)
	}
}

func findAllCerts() map[string][]string {
	cacertsList := map[string][]string{} // to initialize the map to empty map.

	certsList := make([]string, 0)
	err := filepath.Walk("C:\\Users\\Chaitanya.Kolluru\\Desktop\\AA-new\\tuna1\\go\\src\\practice",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error walking..")
				os.Exit(3)
			}
			certsList = append(certsList, path)

			return nil
		})
	if err != nil {
		fmt.Println("Error walking..")
		os.Exit(3)
	}

	for _, file := range certsList {
		fileList := strings.Split(file, "\\")
		switch fileList[len(fileList)-1] {
		case "cacerts":
			cacertsList["cacerts"] = append(cacertsList["cacerts"], file)
		case "trust.p12":
			cacertsList["trust.p12"] = append(cacertsList["trust.p12"], file)
		}

	}

	return cacertsList
}

func parseJSON() map[string][]string {
	cwd, _ := os.Getwd()
	jsonFile := cwd + "\\aa_root_ca_bundle.json"
	data, _ := ioutil.ReadFile(jsonFile)

	var dat map[string][]string
	if err := json.Unmarshal(data, &dat); err != nil {
		fmt.Println("Cannot unmarshal json data")
		os.Exit(2)
	}
	return dat
}

func main() {
	aaRootJSONData := parseJSON()
	cacertsList := findAllCerts()

	for _, caFile := range aaRootJSONData["ca_bundle"] {
		processCaFile(cacertsList, caFile)
	}

}
