// The builder pattern is a creational design pattern that allows you to construct complex objects step by step.
// This can be useful when you want to create objects with many optional parameters,
// or when you want to provide a clean interface for creating objects that belong to a large class hierarchy.
// In Go, you can implement a builder pattern using a separate “builder” type that exposes a fluent interface for constructing the desired object, like this:

// the example given in the post is below, but this can be built better using options-pattern written in the main directory

// type Builder interface {
//     SetOption1(valuel string) Builder
//     SetOption2(value2 int) Builder
//     SetOption3(value bool) Builder
//     BuildO *Product
// }

// type productBuilder struct {
//     option1 string
//     option2 int
//     option3 bool
// }

// func (b *productBuilder) SetOption1(value string) Builder {
//     b.option1 = value
//     return b
// }

// func (b *productBuilder) SetOption2(value int) Builder {
//     b.option2 = value
//     return b
// }

// func (b *productBuilder) SetOption3(value bool) Builder {
//     b.option3 = value
//     return b
// }

// func (b *productBuilder) Build() *Product {
//     return &Product{
//         Option1: b.option1,
//         Option2: b.option2,
//         Option3: b.option3,
//     }
// }

// func NewBuilder() Builder {
//     return &productBuilder{}
// }

package main

import "fmt"

type productBuidler struct {
	option1 string
	option2 int
	option3 bool
}

type buildlerFunc func(p *productBuidler)

func withOption1(o string) buildlerFunc {
	return func(p *productBuidler) { p.option1 = o }
}

func withOption2(o int) buildlerFunc {
	return func(p *productBuidler) { p.option2 = o }
}

func withOption3(o bool) buildlerFunc {
	return func(p *productBuidler) { p.option3 = o }
}

func newProductBuilder(opts ...buildlerFunc) *productBuidler {
	p := &productBuidler{}
	for _, f := range opts {
		f(p)
	}

	return p
}

func main() {
	testP := newProductBuilder(withOption1("chai"), withOption2(1), withOption3(false))
	fmt.Println(*testP)

}
