package terraformfunctions

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//used for init an plan runs
func TerraformRuns(action, cwd string) {

	var stdout, stderr bytes.Buffer

	if err := os.Chdir(cwd + "\\.temp"); err != nil {
		log.Fatal(err)
	}
	execCmd := exec.Command("C:\\Windows\\System32\\terraform-dir\\terraform", action)
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	if err := execCmd.Run(); err != nil {
		fmt.Println(string(stderr.Bytes()))
		os.Exit(1)
	}

	fmt.Println(string(stdout.Bytes()))

	if err := os.Chdir(cwd); err != nil {
		log.Fatal(err)
	}
}

// used for apply runs
func TerraformRunsApply(cwd string) {

	var stdout, stderr bytes.Buffer

	if err := os.Chdir(cwd + "\\.temp"); err != nil {
		log.Fatal(err)
	}
	execCmd := exec.Command("C:\\Windows\\System32\\terraform-dir\\terraform", "apply", "--auto-approve")
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	if err := execCmd.Run(); err != nil {
		fmt.Println(string(stderr.Bytes()))
		os.Exit(1)
	}

	fmt.Println(string(stdout.Bytes()))

	if err := os.Chdir(cwd); err != nil {
		log.Fatal(err)
	}
}
