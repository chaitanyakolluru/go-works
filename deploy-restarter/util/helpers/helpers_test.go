package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertmapToString(t *testing.T) {

	var tests = []struct {
		data map[string]string
		want string
	}{{map[string]string{"key": "value"}, `{"key":"value"}`}}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %s", tt.data, tt.want)
		t.Run(testname, func(t *testing.T) {
			result := ConvertMapToString(tt.data)
			if result != tt.want {
				t.Errorf("ConvertmapToString results %s, want %s", result, tt.want)
			}
		})
	}
}

func TestEncryptConfigmapData(t *testing.T) {
	// encrypted string from running the openssl command
	want := "e43abcf3375244839c012f9633f95862d232a95b00d5bc7348b3098b9fed7f32"

	result, _ := EncryptConfigmapData(map[string]string{"key": "value"})
	if result != want {
		t.Errorf("EncryptConfigmapData() results: %s, want: %s", result, want)
	}
}

func TestStringifyAnnotationMap(t *testing.T) {
	annotation := map[string]string{"key": "value"}
	resultWant := `{"spec":{"template":{"metadata":{"annotations":{"key":"value"}}}}}`
	result, _ := StringifyAnnotationMap(annotation)

	if result != resultWant {
		t.Errorf("stringifyAnnotationMap() results: %s, want: %s", result, resultWant)
	}
}

func TestMatchAnnotationsUpdateAnnotationTrue(t *testing.T) {
	annotation := map[string]string{"heb.com/namespace-configmap": "encrypted data"}
	resultWant := `{"spec":{"template":{"metadata":{"annotations":{"heb.com/namespace-configmap":"encrypted data"}}}}}`
	resultUpdateAnnotation, resultGot, _ := MatchAnnotations(annotation, "encrypted data", "namespace", "configmap")

	if resultUpdateAnnotation {
		if !reflect.DeepEqual(resultGot, resultWant) {
			t.Errorf("MatchAnnotations() failed, result: %v, %v, want: %v, %v", resultUpdateAnnotation, resultGot, true, resultWant)
		}
	}
}

func TestMatchAnnotationsUpdateAnnotationTrueAnnoatationExists(t *testing.T) {
	annotation := map[string]string{"key": "value", "heb.com/namespace-configmap": "some other value"}
	annotationWant := `{"spec":{"template":{"metadata":{"annotations":{"heb.com/namespace-configmap":"encrypted data"}}}}}`
	resultUpdateAnnotation, resultAnnotation, _ := MatchAnnotations(annotation, "encrypted data", "namespace", "configmap")

	if resultUpdateAnnotation {
		if !reflect.DeepEqual(resultAnnotation, annotationWant) {
			t.Errorf("MatchAnnotations() failed, result: %v, %v, want: %v, %v", resultUpdateAnnotation, resultAnnotation, true, annotationWant)
		}
	}
}

func TestMatchAnnotationsUpdateAnnotationTrueAnnoatationDoesntExist(t *testing.T) {
	annotation := map[string]string{"key": "value"}
	annotationWant := `{"spec":{"template":{"metadata":{"annotations":{"heb.com/namespace-configmap":"encrypted data"}}}}}`
	resultUpdateAnnotation, resultAnnotation, _ := MatchAnnotations(annotation, "encrypted data", "namespace", "configmap")

	if resultUpdateAnnotation {
		if !reflect.DeepEqual(resultAnnotation, annotationWant) {
			t.Errorf("MatchAnnotations() failed, result: %v, %v, want: %v, %v", resultUpdateAnnotation, resultAnnotation, true, annotationWant)
		}
	}
}

func TestMatchAnnotationsAnnotationParamNil(t *testing.T) {
	annotationWant := `{"spec":{"template":{"metadata":{"annotations":{"heb.com/namespace-configmap":"encrypted data"}}}}}`
	resultUpdateAnnotation, resultAnnotation, _ := MatchAnnotations(nil, "encrypted data", "namespace", "configmap")

	if resultUpdateAnnotation {
		if !reflect.DeepEqual(resultAnnotation, annotationWant) {
			t.Errorf("MatchAnnotations() failed, result: %v, %v, want: %v, %v", resultUpdateAnnotation, resultAnnotation, true, annotationWant)
		}
	}
}

func TestMatchAnnotationsUpdateAnnotationFalse(t *testing.T) {
	annotation := map[string]string{"key": "value", "heb.com/namespace-configmap": "encrypted data"}
	resultUpdateAnnotation, _, _ := MatchAnnotations(annotation, "encrypted data", "namespace", "configmap")

	if resultUpdateAnnotation {
		t.Errorf("MatchAnnotations() failed, result: %v, _, want: %v, _", resultUpdateAnnotation, false)
	}
}
