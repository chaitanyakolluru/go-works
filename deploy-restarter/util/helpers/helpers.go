package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// helper function to reduce code reuse
func helperUpdateAnnotation(annotationKey, annotationValue string) (bool, string, error) {
	resultAnnotation := map[string]string{annotationKey: annotationValue}
	stringify, err := StringifyAnnotationMap(resultAnnotation)
	return true, stringify, err
}

// return string representation of annotation to be patched into the deployment
func StringifyAnnotationMap(annotationMap map[string]string) (string, error) {

	newMap := make(map[string]interface{})
	newMap["spec"] = make(map[string]interface{})
	newMap["spec"].(map[string]interface{})["template"] = make(map[string]interface{})
	newMap["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"] = make(map[string]interface{})
	newMap["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"].(map[string]interface{})["annotations"] = make(map[string]interface{})

	for k, v := range annotationMap {
		newMap["spec"].(map[string]interface{})["template"].(map[string]interface{})["metadata"].(map[string]interface{})["annotations"].(map[string]interface{})[k] = v
	}

	jsonData, err := json.Marshal(newMap)
	if err != nil {
		return "", fmt.Errorf("cannot json marshal annotation %v, Error: %v", newMap, err.Error())
	}

	return string(jsonData), nil
}

// checks to see if deployment object's config map annotation is different from returned value encryptedData
// and returns a true and the annotation map if so.
func MatchAnnotations(annotations map[string]string, encryptedData string, namespace string, configmap string) (bool, string, error) {
	annotationKey := "heb.com/" + namespace + "-" + configmap
	annotationValue := encryptedData

	if annotations == nil {
		return helperUpdateAnnotation(annotationKey, annotationValue)
	} else {
		if value, exists := annotations[annotationKey]; exists {
			if value != annotationValue {
				return helperUpdateAnnotation(annotationKey, annotationValue)
			}
		} else {
			return helperUpdateAnnotation(annotationKey, annotationValue)
		}
	}

	return false, "", nil
}

// convert dictionary to string
func ConvertMapToString(data map[string]string) string {

	converted, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Cannot marshal config map data. Error: %v", err.Error())
	}

	return string(converted)
}

// get configmap data and create an encrypted string from it.
func EncryptConfigmapData(data map[string]string) (string, error) {
	converted := ConvertMapToString(data)
	encryptedString := ""

	// openssl to encrypt config map data
	cmd := exec.Command("openssl", "dgst", "-sha256")

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("cannot create stdin pipe, Error: %v", err.Error())
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("cannot create stdout pipe, Error: %v", err.Error())
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("cannot start openssl command, Error: %v", err.Error())
	}

	if _, err := stdinPipe.Write([]byte(converted)); err != nil {
		return "", fmt.Errorf("cannot write to stdin pipe, Error: %v", err.Error())
	}

	if err := stdinPipe.Close(); err != nil {
		return "", fmt.Errorf("cannot close stdin pipe, Error: %v", err.Error())
	}

	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		encryptedString += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("cannot scan output from openssl, Error: %v", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("cannot wait for openssl command, Error: %v", err.Error())
	}

	splitted := strings.Split(encryptedString, "=")
	return strings.TrimSpace(splitted[len(splitted)-1]), nil
}
