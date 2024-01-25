package main

import "fmt"

func main() {

	t := map[int]interface{}{1: 2, 3: map[int]int{4: 6}, 4: map[int]int{5: 6}}
	fmt.Println(t)
}
