package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(input []int) []int {
	seen := map[int]bool{}
	newArray := []int{}
	for _, item := range input {
		ok, _ := seen[item]
		if !ok {
			seen[item] = true
			newArray = append(newArray, item)
		} else {
			newArray = append(newArray, 0)
		}

	}

	return newArray
}

func twoBigNodes(input []int) []int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j] // Reverse the comparison
	})
	return input[:2]
}

func recordInputIndexes(input []int) map[int]int {
	recorder := map[int]int{}
	for index, item := range input {
		recorder[item] = index
	}
	return recorder
}
func findIndexesAndCalculateArea(indexRecorder map[int]int, largestNodes []int) int {
	var oneSide int
	if largestNodes[0]-largestNodes[1] == 0 {
		oneSide = largestNodes[0]
	} else {
		oneSide = largestNodes[0] - (largestNodes[0] - largestNodes[1])
	}
	firstNodeIndex := indexRecorder[largestNodes[0]]
	secondNodeIndex := indexRecorder[largestNodes[1]]
	secondSide := firstNodeIndex - secondNodeIndex
	if secondSide < 0 {
		secondSide = -secondSide
	}
	return oneSide * secondSide
}
func maxArea(height []int) int {
	dedupedArray := removeDuplicates(height)
	fmt.Println(dedupedArray)
	indexRecorder := recordInputIndexes(dedupedArray)
	fmt.Print(indexRecorder)
	twoLargestNodes := twoBigNodes(dedupedArray)
	fmt.Print(twoLargestNodes)
	return findIndexesAndCalculateArea(indexRecorder, twoLargestNodes)
}

func main() {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{1, 1}
	fmt.Println(maxArea(height))
}
