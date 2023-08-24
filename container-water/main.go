package main

import (
	"fmt"
)

func calculateCommonY(x, y int) int {
	if x > y {
		return x - subtractAndReturnPositive(x, y)
	} else {
		return y - subtractAndReturnPositive(x, y)
	}
}

func calculateArea(item1, item2 []int) int {
	return subtractAndReturnPositive(item2[0], item1[0]) * calculateCommonY(item1[1], item2[1])
}

func subtractAndReturnPositive(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func maxArea(height []int) (biggestArea int) {
	for index, item := range height {
		iter := 1
		for index+iter < len(height) {
			area := calculateArea([]int{index, item}, []int{index + iter, height[index+iter]})
			if area > biggestArea {
				biggestArea = area
			}
			iter += 1
		}
	}
	return
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
