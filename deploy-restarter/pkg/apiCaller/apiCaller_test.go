package apiCaller

import (
	"context"
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	fakeclient "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

// helper function to create fake deployment objects used to test {Get/Update}DeploymentAnnotations functions
func createFakeDeployment(apiCaller apiCaller, t *testing.T) *v1.Deployment {
	deployment := v1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "deployment", Namespace: "namespace"},
		Spec: v1.DeploymentSpec{Template: corev1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{
			Name:        "config-map",
			Namespace:   "namespace",
			Annotations: map[string]string{"annotation-key": "annotation-value"},
		}}}}

	// add deployment to fake clientset
	deploymentItem, err := apiCaller.clientset.AppsV1().Deployments("namespace").Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create deployment: %v", deployment)
	}

	return deploymentItem
}

// helper function to test CreateClientSet() function and run the function
func setupMocksForCreateClientSet(t *testing.T, restError error, kubernetesError error) (*kubernetes.Clientset, error) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := NewMockRestInterface(ctrl)
	mockRest.EXPECT().InClusterConfig().AnyTimes().Return(&rest.Config{}, restError)

	mockKubernetes := NewMockKubernetesInterface(ctrl)
	mockKubernetes.EXPECT().NewForConfig(&rest.Config{}).AnyTimes().Return(&kubernetes.Clientset{}, kubernetesError)

	result, err := CreateClientSet(mockRest, mockKubernetes)
	return result, err
}

func TestGetConfigMap(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

	// Create the config map object
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "config-map",
			Namespace: "namespace",
		},
		Data: map[string]string{
			"key": "value",
		},
	}

	// Add the config map to the fake clientset
	_, err := apiCaller.clientset.CoreV1().ConfigMaps("namespace").Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create config map: %v", err)
	}

	result, _ := apiCaller.GetConfigMap("config-map", "namespace")
	if !reflect.DeepEqual(result, map[string]string{"key": "value"}) {
		t.Errorf(`GetConfigMap() failed, result: %v, want: map[string]string{"key": "value"}`, result)
	}
}

func TestGetConfigMapGetErrors(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

	_, err := apiCaller.GetConfigMap("config-map", "namespace")

	if err == nil {
		t.Errorf(`GetConfigMap() did not error with log output: Cannot get config map data, Error: configmaps "config-map" not found`)
	}

}

func TestGetDeploymentAnnotations(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

_:
	createFakeDeployment(apiCaller, t)

	resultAnnotation, _ := apiCaller.GetDeploymentAnnotations("deployment", "namespace")

	if !reflect.DeepEqual(resultAnnotation, map[string]string{"annotation-key": "annotation-value"}) {
		t.Errorf(`GetDeploymentAnnotations() failed, resultAnnotation: %v, want: map[string]string{"annotation-key": "annotation-value"}`, resultAnnotation)
	}
}

func TestPatchDeploymentAnnotations(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

	createFakeDeployment(apiCaller, t)

	annotationString := `{"spec": {"template": {"metadata": {"annotations": {"new-annotation-key": "new-annotation-value"}}}}}`
	annotationWant := map[string]string{"new-annotation-key": "new-annotation-value"}

	depItem, _ := apiCaller.PatchDeploymentAnnotations("namespace", "deployment", annotationString)

	if depItem.Spec.Template.Annotations["new-annotation-key"] != annotationWant["new-annotation-key"] {
		t.Errorf(`PatchDeploymentAnnotations() failed, resultAnnotation: %v, want: %v`, depItem.Spec.Template.Annotations["new-annotation-key"], annotationWant["new-annotation-key"])
	}
}

func TestPatchDeploymentAnnotationsUpdateErrors(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

	_, err := apiCaller.PatchDeploymentAnnotations("namespace", "deployment", "")

	if err == nil {
		t.Errorf(`PatchDeploymentAnnotations() did not error. Want: %v`, err.Error())
	}

}

func TestGetDeploymentAnnotationsGetErrors(t *testing.T) {
	apiCaller := apiCaller{clientset: fakeclient.NewSimpleClientset()}

	_, err := apiCaller.GetDeploymentAnnotations("deployment", "namespace")

	if err == nil {
		t.Errorf(`GetDeploymentAnnotations() did not error. Want: %v`, err.Error())
	}

}

func TestCreateClientSet(t *testing.T) {
	result, _ := setupMocksForCreateClientSet(t, nil, nil)

	if reflect.TypeOf(*result) != reflect.ValueOf(&kubernetes.Clientset{}).Elem().Type() {
		t.Errorf("CreateClientSet() returns result of wrong type, Result: %v, want: %v", *result, reflect.ValueOf(&kubernetes.Clientset{}).Elem().Type())
	}
}

func TestCreateClientSetErrorsWithLoadingInClusterConfig(t *testing.T) {

	_, err := setupMocksForCreateClientSet(t, errors.New("Error with rest.InclusterConfig()"), nil)

	if err == nil {
		t.Errorf(`CreateClientSet() did not error with log output: Cannot gather in cluster config, Error: Error with rest.InclusterConfig()`)
	}

}

func TestCreateClientSetErrorsWithK8sNewForConfig(t *testing.T) {

	_, err := setupMocksForCreateClientSet(t, nil, errors.New("Error with kubernetes.NewForConfig()"))

	if err == nil {
		t.Errorf(`CreateClientSet() did not error with log output: Cannot create new clientset, Error: Error with kubernetes.NewForConfig()`)
	}

}
