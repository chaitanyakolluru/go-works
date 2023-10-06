package interfaces

import "fmt"

type laptop struct {
	name, os, ram string
	year          int
	costs         float64
}

type car struct {
	name, make, model string
	year              int
	costs             float64
}

type CostInt interface {
	cost(string) float64
}

func CostIntFunc(c CostInt, ss string) float64 {
	return c.cost(ss)
}

func (c *laptop) cost(ss string) float64 {
	fmt.Println(ss)
	return c.costs
}

func (l *car) cost(pp string) float64 {
	fmt.Println(pp)
	return l.costs
}

func Interfaces() {

	l := &laptop{name: "mac", os: "OSX", ram: "32G", year: 2022, costs: 3000.00}
	c := &car{name: "fo", make: "subaru", model: "forester", year: 2021, costs: 32000.00}

	fmt.Println(CostIntFunc(l, "lap"), CostIntFunc(c, "car"))

}
