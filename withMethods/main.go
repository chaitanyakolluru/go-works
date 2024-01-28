package main

import (
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

func main() {

	s := &Chai{}
	c := &Chai{}
	s.withItem("chaitanya").
		withItem(false).
		withItem(0)

	c.withItemName("chaitanya", "something").
		withItemName(false, "sselse").
		withItemName(0, "thought")

}
