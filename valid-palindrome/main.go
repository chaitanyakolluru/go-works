package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleanLowerString := strings.ToLower(reg.ReplaceAllString(s, ""))

	cleanLowerStringSlice := strings.Split(cleanLowerString, "")
	reverseStringSlice := make([]string, len(cleanLowerStringSlice))

	for i := len(cleanLowerStringSlice) - 1; i >= 0; i-- {
		reverseStringSlice[len(cleanLowerStringSlice)-1-i] = cleanLowerStringSlice[i]
	}

	reverseString := strings.Join(reverseStringSlice, "")
	fmt.Println(reverseString)
	return reverseString == cleanLowerString

}

func main() { fmt.Println("valid") }
