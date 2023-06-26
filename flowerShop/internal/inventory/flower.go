package inventory

import "fmt"

type Flower struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

func CreateFlower(color, name string, no int) *Flower {
	return &Flower{
		Color: fmt.Sprintf("color-%d", no),
		Name:  fmt.Sprintf("name-%d", no),
	}
}

func CreateFlowerInventory(no int) (flowers []*Flower) {
	for i := 0; i < no; i++ {
		flowers = append(flowers, CreateFlower("color", "name", i))
	}
	return
}
