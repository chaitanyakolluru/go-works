package main

import "fmt"

// I want to test if we have a var that has data put inside a for loop or something, if that data persists if it's being
// written to a var vs a pointer.

type SomeItem struct {
	option1 string
	option2 string
	option3 string
}

func main() {
	someVar := SomeItem{}
	somePointer := &SomeItem{}

	for _, set := range [][]string{{"1", "2", "3"}} {
		someVar.option1 = set[0]
		someVar.option2 = set[1]
		someVar.option3 = set[2]

		somePointer.option1 = set[0]
		somePointer.option2 = set[1]
		somePointer.option3 = set[2]
	}

	fmt.Println("var", someVar)
	fmt.Println("pointer", somePointer)

}
