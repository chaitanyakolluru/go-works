package main

import (
	"fmt"
)

func threeSumClosest(nums []int, target int) (closestSum int) {
	for i, v := range nums {
		for j, v2 := range nums {
			if j != i {
				for k, v3 := range nums {
					if k != j && k != i {
						sum := v + v2 + v3
						fmt.Println("sum and all", sum, i, j, k)
						if (target - closestSum) > (target - sum) {
							closestSum = sum
						}
					}
				}
			}
		}
	}

	return
}

func main() {
	fmt.Println("this is 3 sum closest")
}
