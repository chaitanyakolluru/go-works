package terraformfunctions

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"
)

func TestTerraformRuns(t *testing.T) {
	cwd, _ := os.Getwd()
	if err := os.Mkdir(cwd+string(os.PathSeparator)+".temp", 0755); err != nil {
		log.Fatal(err)
	}
	outGot := captureStdout(TerraformRuns)
	outWant := "Terraform initialized in an empty directory"
	if m, _ := regexp.MatchString(outWant, outGot); !m {
		t.Errorf("TerraformRuns output is incorrect, got: %s, want: %s", outGot, outWant)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}

func captureStdout(f func(string, ...string)) string {
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

	f(cwd, "init")

	w.Close()

	return <-out

}
