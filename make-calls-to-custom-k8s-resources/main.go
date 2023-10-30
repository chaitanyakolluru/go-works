package main

import (
	"flag"
	"fmt"
	"log"

	utils "git.heb.com/kub/composition-functions/resources/utils"
	rv1alpha1 "git.heb.com/kub/xrd/api/iam/v1alpha1"
	psjav1alpha1 "git.heb.com/provider-simplejsonapp/apis/records/v1alpha1"
)

func main() {

	typed := flag.String("typed", "records", "usage")
	flag.Parse()

	if *typed == "records" {
		fmt.Println("going records")
		d, err := utils.NewDynamicClient(
			utils.WithGroup("records.simplejsonapp.crossplane.io"),
			utils.WithVersion("v1alpha1"),
			utils.WithResource("records"),
		)

		if err != nil {
			log.Fatalf("error when creating dynamic client, err: %s", err)
		}

		var rr psjav1alpha1.Record
		rr, err = utils.GetCustomResource[psjav1alpha1.Record](d, rr, "example-record-new", "")

		if err != nil {
			log.Fatalf("error with getting custom resource, err: %s", err)
		}

		fmt.Printf("Using Utils: record is: %s, spec.Location: %s\n", rr.ObjectMeta.Name, rr.Spec.ForProvider.Location)
	}

	if *typed == "rosters" {
		fmt.Println("going roster")
		dd, err := utils.NewDynamicClient(
			utils.WithGroup("iam.kub.heb.com"),
			utils.WithVersion("v1alpha1"),
			utils.WithResource("rosters"),
		)

		if err != nil {
			log.Fatalf("error when creating dynamic client, err: %s", err)
		}

		var r rv1alpha1.Roster
		r, err = utils.GetCustomResource[rv1alpha1.Roster](dd, r, "example-roster", "default")

		if err != nil {
			log.Fatalf("error with getting custom resource, err: %s", err)
		}

		fmt.Printf("stuff: %v", r.Spec)
	}

}
