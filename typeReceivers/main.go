package main

import (
	"fmt"
)

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
