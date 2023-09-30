package main

import "strings"

func isLongPressedName(name string, typed string) bool {

	nameSlice := strings.Split(name, "")
	typedSlice := strings.Split(typed, "")

	matchCounter := 0
	for i := 0; i < len(nameSlice); i++ {
		touched := false
		for j := matchCounter; j < len(typedSlice); j++ {
			if nameSlice[i] == typedSlice[j] {
				matchCounter++
				touched = true
				if i < len(nameSlice)-1 && nameSlice[i] == nameSlice[i+1] {
					break
				}
			} else if touched {
				if i == len(nameSlice)-1 {
					for k := j; k < len(typedSlice); k++ {
						if nameSlice[i] != typedSlice[k] {
							return false
						}
					}
				}

				break
			} else {
				return false
			}
		}

		if !touched {
			return false
		}

	}

	return true
}

func main() {}
