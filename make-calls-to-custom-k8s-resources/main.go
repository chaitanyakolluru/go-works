package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	utils "git.heb.com/kub/composition-functions/resources/utils"

	psjav1alpha1 "git.heb.com/provider-simplejsonapp/apis/records/v1alpha1"

	"k8s.io/client-go/rest"
)

func createClient() *dynamic.DynamicClient {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("cannot retrieve in-cluster config, err: %s", err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("cannot create dynamic client, err: %s", err)
	}

	return client
}

func getRecord(client *dynamic.DynamicClient, record string) {
	recordRes := schema.GroupVersionResource{Group: "records.simplejsonapp.crossplane.io", Version: "v1alpha1", Resource: "records"}

	result, err := client.Resource(recordRes).Get(context.TODO(), record, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	var r psjav1alpha1.Record
	utils.Hydrate(result, &r)

	fmt.Printf("record is: %s, spec.Location: %s", r.ObjectMeta.Name, r.Spec.ForProvider.Location)
}

func main() {
	getRecord(createClient(), "example-record-new")
}
