package main

import (
	"fmt"
)

func main() {
	inpArr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	n := len(inpArr)

	outArrCoor := make([][][]int, 0)
	for n > 0 {
		itemO := make([][]int, 0)
		for i := 0; i < len(inpArr); i++ {
			itemOI := []int{}
			itemOI = append(itemOI, indexOf(inpArr, inpArr[len(inpArr)-1])-i)
			itemOI = append(itemOI, (len(inpArr)-1)-(n-1))
			itemO = append(itemO, itemOI)

		}
		outArrCoor = append(outArrCoor, itemO)
		n -= 1
	}

	//fmt.Println(outArrCoor)

	outArr := make([][]int, 0)
	for _, itemOO := range outArrCoor {
		itemOOA := make([]int, 0)
		for _, iOO := range itemOO {
			itemOOA = append(itemOOA, inpArr[iOO[0]][iOO[1]])
		}

		outArr = append(outArr, itemOOA)
	}

	fmt.Println(outArr)

}

func indexOf(data [3][3]int, element [3]int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
