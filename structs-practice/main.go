package main

import (
	"fmt"
	"structs-practice/interfaces"
)

type Salary struct {
	basic, hra, ta float64
}

type Employee struct {
	name   string
	age    int
	salary Salary
}

func (e *Employee) info() {
	fmt.Println(e.name)
	fmt.Println(e.age)
	fmt.Println(e.salary)
}

type infor interface {
	info()
}

func main() {

	// var e Employee
	// e.name = "chai"
	// e.age = 34
	// e.salary = Salary{100, 100.04, 100.00}

	// e.info()

	// var i infor
	// i = &e
	// i.info()

	interfaces.Interfaces()

}
