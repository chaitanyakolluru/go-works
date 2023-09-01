package main

import (
	"fmt"
)

func calculateCommonY(x, y int) int {
	return y - x
}

func calculateArea(item1, item2 []int) int {
	return (item2[0] - item1[0]) * calculateCommonY(item1[1], item2[1])
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
