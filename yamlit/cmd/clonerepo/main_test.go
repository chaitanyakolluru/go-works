package clonerepo

import (
	"os"
	"testing"
)

func TestCloneRepo(t *testing.T) {
	cwd, _ := os.Getwd()

	CloneRepo(cwd, "", "")
	if _, err := os.Stat(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform" + string(os.PathSeparator) + ".git"); os.IsNotExist(err) {
		if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
			t.Fatal(err)
		}
		t.Errorf("Clonerepo function didn't clone AA/Terraform repo as expected.")
	}
	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}
