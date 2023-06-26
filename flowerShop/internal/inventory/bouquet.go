package inventory

import "fmt"

type Bouquet struct {
	Name    string    `json:"name"`
	Flowers []*Flower `json:"flowers"`
}

func CreateBouquet(name string, flowers []*Flower, no int) *Bouquet {
	return &Bouquet{Name: fmt.Sprintf("Name-%d", no), Flowers: flowers}
}

func CreateBouquetInventory(no int) (bouquets []*Bouquet) {
	for i := 0; i < no; i++ {
		bouquets = append(bouquets, CreateBouquet("size", []*Flower{}, i))
	}
	return
}
