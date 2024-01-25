package terraformfunctions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

//TerraformRuns used for init an plan runs
func TerraformRuns(cwd string, action ...string) {

	if err := os.Chdir(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		log.Fatal(err)
	}
	execCmd := exec.Command("terraform", action...)

	execOutReader, err := execCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(execOutReader)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		wg.Done()
	}()
	if scanner.Err() != nil {
		log.Fatal(err)
	}
	stderr, err := execCmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := execCmd.Start(); err != nil {
		log.Fatal(err)
	}

	stderrStr, _ := ioutil.ReadAll(stderr)
	if string(stderrStr) != "" {
		fmt.Println(string(stderrStr))
		os.Exit(1)
	}

	wg.Wait()
	if err := execCmd.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := os.Chdir(cwd); err != nil {
		log.Fatal(err)
	}
}
