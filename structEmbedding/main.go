package main

import "fmt"

type goodThing struct {
	kindOfCool string
}

type something struct {
	goodThing
}

func main() {

	ss := something{
		goodThing: goodThing{
			kindOfCool: "chaitanya",
		},
	}

	fmt.Println(ss, ss.kindOfCool)

}
