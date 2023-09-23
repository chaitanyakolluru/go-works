package main

import (
	"fmt"
	"strings"
)

func increaseCount(openParens []string, i string, count int) ([]string, int) {
	openParens = append(openParens, i)
	count++
	return openParens, count
}

func decreaseCount(openParens []string, paren []string, count int) ([]string, int, bool) {
	if len(openParens) >= 1 {
		if paren[0] == openParens[len(openParens)-1] {
			count--
			openParens = openParens[:len(openParens)-1]
		} else {
			return openParens, count, false
		}
	} else {
		return openParens, count, false
	}
	return openParens, count, true
}

func isValid(s string) bool {
	paren1 := []string{"(", ")"}
	paren2 := []string{"[", "]"}
	paren3 := []string{"{", "}"}
	listOfParens := [][]string{paren1, paren2, paren3}
	var paren1Count, paren2Count, paren3Count int
	var ret bool
	openParens := make([]string, 0)

	for _, i := range strings.Split(s, "") {
		for k, li := range listOfParens {
			if i == li[0] {
				switch k {
				case 0:
					openParens, paren1Count = increaseCount(openParens, i, paren1Count)
				case 1:
					openParens, paren2Count = increaseCount(openParens, i, paren2Count)
				case 2:
					openParens, paren3Count = increaseCount(openParens, i, paren3Count)
				}
				break
			}
			if i == li[1] {
				switch k {
				case 0:
					openParens, paren1Count, ret = decreaseCount(openParens, paren1, paren1Count)
					if !ret {
						return false
					}
				case 1:
					openParens, paren2Count, ret = decreaseCount(openParens, paren2, paren2Count)
					if !ret {
						return false
					}
				case 2:
					openParens, paren3Count, ret = decreaseCount(openParens, paren3, paren3Count)
					if !ret {
						return false
					}
				}
				break
			}
		}
	}

	return paren1Count == 0 && paren2Count == 0 && paren3Count == 0
}

func isValidFromChatGPT(s string) bool {
	parenMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != parenMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	//fmt.Println(isValid("()[]{}"))
	//	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
