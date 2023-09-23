package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	paren1 := []string{"(", ")"}
	paren2 := []string{"[", "]"}
	paren3 := []string{"{", "}"}
	listOfParens := [][]string{paren1, paren2, paren3}
	var paren1Count, paren2Count, paren3Count int
	openParens := make([]string, 0)

	for _, i := range strings.Split(s, "") {
		for k, li := range listOfParens {
			if i == li[0] {
				switch k {
				case 0:
					openParens = append(openParens, i)
					paren1Count++
				case 1:
					openParens = append(openParens, i)
					paren2Count++
				case 2:
					openParens = append(openParens, i)
					paren3Count++
				}
			}
			if i == li[1] {
				switch k {
				case 0:
					if len(openParens) >= 1 {
						if paren1[0] == openParens[len(openParens)-1] {
							paren1Count--
							openParens = openParens[:len(openParens)-1]
						} else {
							return false
						}
					} else {
						return false
					}
				case 1:
					if len(openParens) >= 1 {
						if paren2[0] == openParens[len(openParens)-1] {
							paren2Count--
							openParens = openParens[:len(openParens)-1]
						} else {
							return false
						}
					} else {
						return false
					}

				case 2:
					if len(openParens) >= 1 {
						if paren3[0] == openParens[len(openParens)-1] {
							paren3Count--
							openParens = openParens[:len(openParens)-1]
						} else {
							return false
						}
					} else {
						return false
					}

				}
			}
		}
	}

	return paren1Count == 0 && paren2Count == 0 && paren3Count == 0
}

func main() {
	//fmt.Println(isValid("()[]{}"))
	//	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
