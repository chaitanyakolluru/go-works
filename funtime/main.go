package main

import "fmt"

type Different struct {
	name string
	age  int
}

type withDifferentOptions func(d *Different)

func NewDifferent(opts ...withDifferentOptions) *Different {
	d := &Different{}

	for _, v := range opts {
		v(d)
	}

	return d
}

func withName(name string) withDifferentOptions {
	return func(d *Different) {
		d.name = name
	}
}

func withAge(age int) withDifferentOptions {
	return func(d *Different) {
		d.age = age
	}
}

type Same struct {
	Different
}

func main() {
	s := Same{
		Different: *NewDifferent(
			withName("chaitanya"),
			withAge(35),
		),
	}

	fmt.Println(s)
}
