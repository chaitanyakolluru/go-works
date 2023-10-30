package main

import (
	"fmt"
	"log"

	utils "git.heb.com/kub/composition-functions/resources/utils"
	rv1alpha1 "git.heb.com/kub/xrd/api/iam/v1alpha1"
)

func main() {

	d, err := utils.NewDynamicClient(
		utils.WithGroup("iam.kub.heb.com"),
		utils.WithVersion("v1alpha1"),
		utils.WithResource("rosters"),
	)

	if err != nil {
		log.Fatalf("error when creating dynamic client, err: %s", err)
	}

	var r rv1alpha1.Roster
	r, err = utils.GetCustomResource[rv1alpha1.Roster](d, r, "example-record-new")

	if err != nil {
		log.Fatalf("error with getting custom resource, err: %s", err)
	}

	fmt.Printf("stuff: %v", r.Spec)
}
