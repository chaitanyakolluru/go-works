// The adapter pattern is a structural design pattern that allows you to adapt one interface to another.
//This can be useful when you have an existing class that provides a certain functionality, but you need to use it in a way that is incompatible with its current interface.
//In Go, you can implement an adapter pattern by creating a wrapper type that implements the desired interface and delegates method calls to the underlying type, like this:

package main

import "fmt"

type Adaptee struct{}

func (a *Adaptee) MakeWhiteNoise() string { return "white noise!!" }

type Adapter struct {
	Adaptee
}

func (aa *Adapter) MakeNoise() string {
	return aa.MakeWhiteNoise()
}

type Target interface {
	MakeNoise() string
}

func NoiseGenerator(t Target) string {
	return t.MakeNoise()
}

func main() {
	fmt.Println("adapter-pattern")
	fmt.Println(NoiseGenerator(&Adapter{}))
}
