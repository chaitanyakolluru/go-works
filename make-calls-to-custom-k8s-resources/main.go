package main

import (
	"fmt"
	"log"

	utils "git.heb.com/kub/composition-functions/resources/utils"
	psjav1alpha1 "git.heb.com/provider-simplejsonapp/apis/records/v1alpha1"
)

func main() {
	d, err := utils.NewDynamicClient(
		utils.WithGroup("records.simplejsonapp.crossplane.io"),
		utils.WithVersion("v1alpha1"),
		utils.WithResource("records"),
	)

	if err != nil {
		log.Fatalf("error when creating dynamic client, err: %s", err)
	}

	result, err := utils.GetCustomResource(d, "example-record-new")

	if err != nil {
		log.Fatalf("error with getting custom resource, err: %s", err)
	}

	var r psjav1alpha1.Record
	utils.Hydrate(result, &r)

	fmt.Printf("Using Utils: record is: %s, spec.Location: %s\n", r.ObjectMeta.Name, r.Spec.ForProvider.Location)
}
