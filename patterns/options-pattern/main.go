package main

import "fmt"

type example struct {
	option1 int
	option2 string
}

type exampleOptions func(e *example)

func withOption1(i int) exampleOptions {
	return func(e *example) {
		e.option1 = i
	}
}

func withOption2(s string) exampleOptions {
	return func(e *example) {
		e.option2 = s
	}
}

func newExample(opts ...exampleOptions) *example {
	e := &example{option1: 1, option2: "default"}

	for _, f := range opts {
		f(e)
	}
	return e
}

func main() {
	example1 := newExample()
	example2 := newExample(withOption1(2), withOption2("value set by options"))

	fmt.Println(*example1, *example2)

	example3 := newExample()
	example3.
		withOption1(3).
		withOption2("value set by alternate options")

	fmt.Println(*example3)
}

// more info on options pattern here: https://michalzalecki.com/golang-options-pattern/#:~:text=Functional%20Options%20Pattern%20also%20called,options%2C%20thus%20the%20pattern%20name.

// another way to implement options with objects.

func (e *example) withOption1(i int) *example {
	e.option1 = i
	return e
}

func (e *example) withOption2(s string) *example {
	e.option2 = s
	return e
}
