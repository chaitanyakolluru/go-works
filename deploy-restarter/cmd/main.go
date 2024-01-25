package main

import (
	"deploy-restarter/pkg/apiCaller"
	"deploy-restarter/util/helpers"
	"flag"
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/util/retry"
)

// prints usage
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: deploy-restarter [options]\n")
	fmt.Fprintf(os.Stderr, "Options: \n")
	flag.PrintDefaults()
}

// validate flags and exit with usage details if validation fails
func validateFlags(configMap string, deployment string, namespace string) {
	printUsage := false
	if deployment == "" {
		fmt.Fprint(os.Stderr, "Flag --configmap must be specified\n\n")
		printUsage = true
	}

	if deployment == "" {
		fmt.Fprint(os.Stderr, "Flag --deployment must be specified\n\n")
		printUsage = true
	}

	if namespace == "" {
		fmt.Fprint(os.Stderr, "Flag --namespace must be specified\n\n")
		printUsage = true
	}

	if printUsage {
		usage()
		os.Exit(1)
	}
}

func main() {

	configMapPtr := flag.String("configmap", "", "changes to this config map cause updated annotation to deployment")
	deploymentPtr := flag.String("deployment", "", "deployment that needs to be updated")
	namespacePtr := flag.String("namespace", "", "namespace under which these resources exist")

	flag.Usage = usage

	flag.Parse()

	configmap := *configMapPtr
	deployment := *deploymentPtr
	namespace := *namespacePtr

	validateFlags(configmap, deployment, namespace)

	apiCaller, err := apiCaller.CreateApiCaller()
	if err != nil {
		log.Fatal(err)
	}

	data, err := apiCaller.GetConfigMap(configmap, namespace)
	if err != nil {
		log.Fatal(err)
	}
	encryptedData, err := helpers.EncryptConfigmapData(data)
	if err != nil {
		log.Fatal(err)
	}

	err = retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		annotations, err := apiCaller.GetDeploymentAnnotations(deployment, namespace)
		if err != nil {
			return err
		}

		updateAnnotation, annotationToPatch, err := helpers.MatchAnnotations(annotations, encryptedData, namespace, configmap)
		if err != nil {
			log.Fatal(err)
		}

		if updateAnnotation {
			_, err := apiCaller.PatchDeploymentAnnotations(namespace, deployment, annotationToPatch)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Cannot get or patch deployment annotation. Error: %v", err.Error())
	}

}
