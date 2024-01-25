package main

import "fmt"

func main() {
	fmt.Println("Simple string") //interpreted string - interpretes as expected.
	fmt.Println(`
	this is multi-line
	string. that can also contain "quotes"`) // raw string - dumps everything as is
	fmt.Println(">")
	fmt.Println("\u2272")
	fmt.Println("LA") // cant use single quotes. rune = string // fmt formatter.
	fmt.Println('L') // here it gives u the number for rune L
}
