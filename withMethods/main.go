package main

import (
	"fmt"
	"reflect"
)

type Chai struct {
	something string
	sselse    bool
	thought   int
}

func (c *Chai) withItemName(i interface{}, varName string) *Chai {
	switch varName {
	case "something":
		c.something = i.(string)
	case "sselse":
		c.sselse = i.(bool)
	case "thought":
		c.thought = i.(int)
	}

	return c
}

func (c *Chai) withItem(i interface{}) *Chai {
	switch reflect.TypeOf(i) {
	case reflect.TypeOf(""):
		c.something = i.(string)
	case reflect.TypeOf(false):
		c.sselse = i.(bool)
	case reflect.TypeOf(0):
		c.thought = i.(int)
	}
	return c
}

type withOption func(c *Chai)

func NewChai(opts ...withOption) *Chai {
	s := &Chai{}

	for _, f := range opts {
		f(s)
	}
	return s
}

func (c *Chai) withFuncOptions(i interface{}, varName string) withOption {
	switch varName {
	case "something":
		return func(c *Chai) { c.something = i.(string) }
	case "sselse":
		return func(c *Chai) { c.sselse = i.(bool) }
	case "thought":
		return func(c *Chai) { c.thought = i.(int) }
	}

	return func(c *Chai) {}
}

func main() {

	s := &Chai{}
	c := &Chai{}
	s.withItem("chaitanya").
		withItem(false).
		withItem(0)

	c.withItemName("chaitanya", "something").
		withItemName(false, "sselse").
		withItemName(0, "thought")

	n := NewChai(
		c.withFuncOptions("chaitanya", "something"),
		c.withFuncOptions(false, "sselse"),
		c.withFuncOptions(0, "thought"),
	)

	fmt.Println(n)

}
