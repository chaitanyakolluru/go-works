package main

import (
	"fmt"
	"strings"
)

type Timepass struct {
	Item  int
	Value int
}

func NewTimepass(i int) *Timepass {
	return &Timepass{
		Item: i,
	}
}

func (t *Timepass) calculateValue(v int) *Timepass {
	t.Value = v * v

	return t
}

func main() {
	t := NewTimepass(5).
		calculateValue(7)

	fmt.Println(t.Value)

	val := (strings.ReplaceAll("platform_registries_-/.", "-", "_"))
	fmt.Println(strings.ReplaceAll(val, ".", "_"))

}
