package main

import (
	"fmt"
	"strings"
)

type Object struct {
	name string
}

type Shape interface {
	Area() int
	WhatAmI() string
}

type Square struct {
	Object
	side int
}

func (s *Square) Area() int {
	return s.side * s.side
}

func (s *Square) WhatAmI() string {
	return s.Object.name
}

type Triangle struct {
	Object
	length int
	breath int
	height int
}

func (t *Triangle) WhatAmI() string {
	return t.Object.name
}

func CreateShapes(size, height, length, breath int) []Shape {
	return []Shape{
		&Square{
			side: size,
			Object: Object{
				name: fmt.Sprintf("square with size %d", size),
			},
		},
		&Triangle{
			length: length,
			height: height,
			breath: breath,
			Object: Object{
				name: fmt.Sprintf("triangle with length %d breath %d height %d", 1, 2, 5),
			},
		}}

}

func (t *Triangle) Area() int {
	return t.length * t.breath * t.height
}

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	return len(strings.Split(words[len(words)-1], ""))
}

func main() {

	shapes := CreateShapes(5, 1, 2, 5)

	for _, shape := range shapes {
		fmt.Println(shape.WhatAmI())
		fmt.Println(shape.Area())
	}
}
