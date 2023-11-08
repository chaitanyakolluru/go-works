package main

import "fmt"

type example struct {
	name     string
	age      int
	location string
}

func (e *example) withName(name string) *example {
	e.name = name
	return e
}

func (e *example) withAge(age int) *example {
	e.age = age
	return e
}

func (e *example) withLocation(location string) *example {
	e.location = location
	return e
}

func main() {
	newE := example{}
	newE.
		withName("chaitanya").
		withAge(22).
		withLocation("austin")

	fmt.Println(newE)
}
