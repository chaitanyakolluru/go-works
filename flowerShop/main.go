package main

import (
	"flowerShop/internal/inventory"
	"fmt"
)

func main() {
	flowerInventory := inventory.CreateFlowerInventory(30)
	bouquetInventory := inventory.CreateBouquetInventory(10)
	bouqets := inventory.PrepareBouquets(flowerInventory, bouquetInventory)

	for b := 0; b < len(bouqets); b++ {
		fmt.Println("bouquet", bouqets[b].Name)
		fmt.Println("flowers")
		for i := 0; i < len(bouqets[b].Flowers); i++ {
			fmt.Println(bouqets[b].Flowers[i].Name, bouqets[b].Flowers[i].Color)
		}
	}
}
