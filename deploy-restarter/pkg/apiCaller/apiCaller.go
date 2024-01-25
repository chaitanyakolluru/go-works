package apiCaller

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// interface for mocks
type RestInterface interface{ InClusterConfig() (*rest.Config, error) }
type KubernetesInterface interface {
	NewForConfig(c *rest.Config) (*kubernetes.Clientset, error)
}

type RestInterfaceImplementer struct{}

func (RestInterfaceImplementer) InClusterConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
}

type KubernetesInterfaceImplementer struct{}

func (KubernetesInterfaceImplementer) NewForConfig(c *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(c)
}

type apiCaller struct{ clientset kubernetes.Interface }

// get configmap data from given configmap
func (a *apiCaller) GetConfigMap(configMap string, namespace string) (map[string]string, error) {
	configmap, err := a.clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMap, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("cannot get config map data, Error: %v", err)
	}

	return configmap.Data, nil
}

// patches given deployment with updated annoatation
func (a *apiCaller) PatchDeploymentAnnotations(namespace string, deployment string, annotation string) (*v1.Deployment, error) {
	depItem, err := a.clientset.AppsV1().Deployments(namespace).Patch(context.TODO(), deployment, types.StrategicMergePatchType, []byte(annotation), metav1.PatchOptions{})
	if err != nil {
		return nil, err
	}

	return depItem, nil
}

// gets deployment.spec.template.metadata.annotations for a given deployment.
func (a *apiCaller) GetDeploymentAnnotations(deployment string, namespace string) (map[string]string, error) {
	deploymentItem, err := a.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deploymentItem.Spec.Template.Annotations, nil
}

// create clientset using pod's k8s data
func CreateClientSet(r RestInterface, k KubernetesInterface) (*kubernetes.Clientset, error) {
	// create clientset
	config, err := r.InClusterConfig()

	if err != nil {
		return nil, fmt.Errorf("cannot gather in cluster config, Error: %v", err.Error())
	}

	clientset, err := k.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("cannot create new clientset, Error: %v", err.Error())
	}

	return clientset, nil
}

// constructor for apiCaller which injects clientset, to be used to make api calls to the apiserver
func CreateApiCaller() (apiCaller, error) {
	clientset, err := CreateClientSet(RestInterfaceImplementer{}, KubernetesInterfaceImplementer{})
	if err != nil {
		return apiCaller{}, fmt.Errorf("cannot create new apicaller, Error: %v", err.Error())
	}
	return apiCaller{clientset: clientset}, nil
}
