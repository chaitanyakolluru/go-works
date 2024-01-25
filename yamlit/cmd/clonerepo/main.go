package clonerepo

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// CloneRepo function clones the repo under cwd/.temp directory
func CloneRepo(cwd, username, token string) {
	username, token = GenerateAuthDetails(username, token)

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(cwd+string(os.PathSeparator)+".temp", 0755); err != nil {
		log.Fatal(err)
	}

	if _, err := git.PlainClone(cwd+string(os.PathSeparator)+".temp"+string(os.PathSeparator)+"terraform", false, &git.CloneOptions{
		URL: "https://ghe.aa.com/AA/terraform.git",
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		},
	}); err != nil {
		log.Fatal(err)
	}
}

// GenerateAuthDetails function extracts username and token if either one of them are empty strings
func GenerateAuthDetails(username, token string) (string, string) {
	errFatal := ""
	if username == "" {
		if usernameTemp, ok := os.LookupEnv("GHE_USERNAME"); !ok {
			errFatal = "--username is not provided and GHE_USERNAME environment variable is not set; "
		} else {
			username = usernameTemp
		}
	}
	if token == "" {
		if tokenTemp, ok := os.LookupEnv("GHE_TOKEN"); !ok {
			errFatal += "--token is not provided and GHE_TOKEN environment variable is not set"
		} else {
			token = tokenTemp
		}
	}

	if errFatal != "" {
		log.Fatal(errFatal)
	}
	return username, token
}
