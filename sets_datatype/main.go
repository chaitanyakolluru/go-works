package main

import "fmt"

func main() {
	type void struct{}
	var member void // creating a member obecjt off of the empty struct void.
	a := make(map[string]void)
	b := make(map[string]void)

	a["a"] = member
	a["b"] = member
	b["a"] = member

	for i := range a {
		if _, ok := b[i]; !ok {
			fmt.Println(i)
		}
	}
}
