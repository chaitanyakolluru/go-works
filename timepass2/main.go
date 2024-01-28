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

type goodThing struct {
	kindOfCool string
}

type something struct {
	goodThing
}

type testReceivers []string

func (t *testReceivers) printAll() {
	for _, v := range *t {
		fmt.Printf("element is: %s\n", v)
	}
}

func (t *testReceivers) addElement(item string) {
	*t = append(*t, item)
}

func (t *testReceivers) removeElement(item string) {
	arr := *t
	for i, v := range arr {
		if v == item {
			arr = append(arr[0:i], arr[i+1:]...)
			break
		}
	}

	*t = arr

}

func main() {

	// s := &Chai{}
	// c := &Chai{}
	// s.withItem("chaitanya").
	//
	//	withItem(false).
	//	withItem(0)
	//
	// c.withItemName("chaitanya", "something").
	//
	//	withItemName(false, "sselse").
	//	withItemName(0, "thought")
	//
	// fmt.Println(*s, *c)
	//
	//	ss := something{
	//		goodThing: goodThing{
	//			kindOfCool: "chaitanya",
	//		},
	//	}
	//
	// fmt.Println(ss, ss.kindOfCool)

	trvar := &testReceivers{}
	trvar.addElement("chaitanya")
	trvar.addElement("is")
	trvar.addElement("a")
	trvar.addElement("good")
	trvar.addElement("person")

	trvar.printAll()

	trvar.removeElement("person")

	trvar.printAll()
}
