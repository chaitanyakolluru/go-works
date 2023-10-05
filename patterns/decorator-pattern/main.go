// The decorator pattern is a structural design pattern that allows you to add new behavior to an existing object dynamically, by wrapping it in a decorator object.
// This can be useful when you want to add new functionality to an object without subclassing it,
// or when you want to add behavior that is conditional on the runtime environment.
// In Go, you can implement a decorator pattern using an interface to represent the base object,
// and a separate “decorator” type that wraps the base object and adds new behavior, like this:
package main

import "fmt"

type Component interface {
	PrintSomething() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) PrintSomething() string {
	return "printing concrete component"
}

type Decorator struct {
	component Component
}

func (d *Decorator) PrintSomething() string {
	return "printing decorator"
}

func main() {
	comp := &ConcreteComponent{}
	fmt.Println(comp.PrintSomething())

	dec := &Decorator{component: comp}
	fmt.Println(dec.component.PrintSomething())
	fmt.Println(dec.PrintSomething())
}
