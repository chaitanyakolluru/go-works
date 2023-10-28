package main

import "fmt"

func main() {

	a := []int{1, 2}
	b := []int{3, 4}

	for _, ii := range append(a, b...) {
		fmt.Println(ii)
	}
}
