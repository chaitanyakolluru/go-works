package main

import (
	"fmt"
	"sort"
)

type difficult struct {
	name string
	age  int
	sex  string
}

func main() {
	data := []difficult{
		{name: "chai", age: 34, sex: "M"},
		{name: "ss", age: 10, sex: "F"},
		{name: "gggg", age: 55, sex: "M"},
	}

	sort.Slice(data, func(i, j int) bool { return data[i].age < data[j].age })

	fmt.Println(data)
}
