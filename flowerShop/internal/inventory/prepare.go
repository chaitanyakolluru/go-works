package inventory

func PrepareBouquets(flowerInventory []*Flower, bouquetInventory []*Bouquet) []*Bouquet {
	flowersleft := len(flowerInventory)
	for b := 0; b < len(bouquetInventory); b++ {
		if flowersleft > 5 {
			flowersForBouquet := flowerInventory[:5]
			flowerInventory = flowerInventory[5:]
			bouquetInventory[b].Flowers = append(bouquetInventory[b].Flowers, flowersForBouquet...)
			flowersleft = len(flowerInventory)
		} else {
			flowersForBouquet := flowerInventory
			flowersForBouquet = append(flowersForBouquet, CreateFlowerInventory(5-flowersleft)...)
			flowerInventory = []*Flower{}
			bouquetInventory[b].Flowers = append(bouquetInventory[b].Flowers, flowersForBouquet...)
			flowersleft = len(flowerInventory)
		}
	}
	return bouquetInventory
}
