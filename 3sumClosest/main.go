package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	closestSum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(target-sum) < abs(target-closestSum) {
					closestSum = sum
				}
			}
		}
	}
	return closestSum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	fmt.Println("this is 3 sum closest")
}
