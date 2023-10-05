// The factory pattern is a creational design pattern that provides an interface for creating objects in a super class, but allows subclasses to alter the type of objects that will be created.
// This can be useful when you want to create objects that belong to a particular family, but you don’t want to specify the exact class of object that will be created.
//In Go, you can implement a factory pattern using a function that returns an interface type, like this:

package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	MakeNoise() string
}

type Dog struct{}
type Cat struct{}

func (d *Dog) MakeNoise() string {
	return "woof!!"
}

func (c *Cat) MakeNoise() string {
	return "meow!!"
}

func NewAnimal(t string) Animal {

	switch t {
	case "woof!!":
		return &Dog{}
	case "meow!!":
		return &Cat{}
	default:
		return nil
	}
}

func main() {
	d := NewAnimal("woof!!")
	c := NewAnimal("meow!!")

	fmt.Println(reflect.TypeOf(d), reflect.TypeOf(c))
}
