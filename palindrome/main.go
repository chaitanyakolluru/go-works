package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {

	digitsAsStrings := strings.Split(strconv.FormatInt(int64(x), 10), "")
	reversedDigitsAsStrings := make([]string, len(digitsAsStrings))

	for i := len(digitsAsStrings) - 1; i >= 0; i-- {
		reversedDigitsAsStrings[len(digitsAsStrings)-1-i] = digitsAsStrings[i]
	}

	reversedFormedString := strings.Join(reversedDigitsAsStrings, "")
	reversedX, _ := strconv.Atoi(reversedFormedString)

	return x == reversedX
}

func main() { isPalindrome(121) }
