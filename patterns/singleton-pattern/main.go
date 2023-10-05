// The singleton pattern is used to ensure that a class has only one instance, and to provide a global access point to that instance.
// This can be useful when you need to manage a shared resource, such as a database connection or configuration object. In Go, you can implement a singleton using the “sync” package to synchronise access to the shared instance, like this:

package main

import (
	"fmt"
	"sync"
)

type Single struct {
	// this could be parameters that represent a singleton resource like a database or something
	id   int
	name string
}

var instance *Single
var once sync.Once

func getSingleInstance(i int, n string) *Single {
	once.Do(func() {
		instance = &Single{id: i, name: n}
	})
	return instance
}

func main() {
	fmt.Println("singleton-pattern")
	instance1 := getSingleInstance(1, "name1")
	instance2 := getSingleInstance(2, "name2")

	fmt.Println(*instance1, *instance2)
}
