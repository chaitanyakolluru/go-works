package main

import "fmt"

var kind = "global var"

func main() {
	xItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	yItems := []int{2, 5, 3, 4, 6, 7, 8, 1, 1, 2}

	for x := range xItems {
		for y := range yItems {
			fmt.Println(x, y, func() bool {
				fmt.Println(kind)
				return x < y
			}())
		}
	}

}
