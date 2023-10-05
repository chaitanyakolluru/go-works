// The bridge pattern is a structural design pattern that allows you to decouple an abstraction from its implementation, so that the two can vary independently.
// This can be useful when you want to change the implementation of an object without affecting its clients,
// or when you want to avoid exposing the implementation details of an object to its clients.
// In Go, you can implement a bridge pattern using an interface to represent the abstraction,
// and a separate “implementation” type to represent the concrete implementation, like this:

// this is basically composition and also making the composed an interface so multiple structs can satisy the composed attribute; constructing an example of my own

package main

import "fmt"

type Dog struct{}

func (d *Dog) MakeNoise() { fmt.Println("bark") }

type Cat struct{}

func (c *Cat) MakeNoise() { fmt.Println("meow") }

type Lion struct{}

func (l *Lion) MakeNoise() { fmt.Println("roar") }

type AnimalsInterface interface {
	MakeNoise()
}

type Animal struct {
	noise AnimalsInterface
}

func (a *Animal) MakeNoise() {
	a.noise.MakeNoise()
}

func main() {
	for _, a := range []*Animal{{noise: &Dog{}}, {noise: &Cat{}}, {noise: &Lion{}}} {
		a.MakeNoise()
	}

}
