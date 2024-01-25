package clonerepo

import (
	"log"
	"os"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/transport/http"
)

func CloneRepo(cwd string) {

	if err := os.RemoveAll(cwd + "\\.temp"); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(cwd+"\\.temp", 0644); err != nil {
		log.Fatal(err)
	}

	if _, err := git.PlainClone(cwd+"\\.temp\\terraform", false, &git.CloneOptions{
		URL: "https://ghe.aa.com/AA/terraform.git",
		Auth: &http.BasicAuth{
			Username: "852047",
			Password: "75eab5d53c8d56d398efba89637fc32f53284ab7",
		},
		Progress: os.Stdout,
	}); err != nil {
		log.Fatal(err)
	}
}
