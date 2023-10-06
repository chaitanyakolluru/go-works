package main

import "fmt"

type Item struct {
	option1 int
	option2 bool
}

type optionFunc func(i *Item)

func withOption1(o int) optionFunc {
	return func(i *Item) { i.option1 = o }
}

func withOption2(o bool) optionFunc {
	return func(i *Item) { i.option2 = o }
}

func NewItem(opts ...optionFunc) *Item {
	i := &Item{}
	for _, f := range opts {
		f(i)
	}
	return i
}

func main() {
	i := NewItem(withOption1(1), withOption2(false))
	fmt.Println(i)
}
