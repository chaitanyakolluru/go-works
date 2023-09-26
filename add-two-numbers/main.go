package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverseAndGetNumber(ll *ListNode) int64 {
	listOfDigits := make([]int, 0)

	tempNode := ll
	for {
		listOfDigits = append(listOfDigits, tempNode.Val)

		if tempNode.Next != nil {
			tempNode = tempNode.Next
		} else {
			break
		}
	}

	for i, j := 0, len(listOfDigits)-1; i < j; i, j = i+1, j-1 {
		listOfDigits[i], listOfDigits[j] = listOfDigits[j], listOfDigits[i]
	}

	var number float64
	numberBig := big.NewFloat(number)
	for i := 0; i <= len(listOfDigits)-1; i++ {
		tempBig := big.NewFloat(float64(listOfDigits[i]))
		numberBig = numberBig.Add(numberBig, tempBig.Mul(tempBig, big.NewFloat(math.Pow(float64(10), float64(len(listOfDigits)-1-i)))))
	}

	result, _ := numberBig.Int64()
	return result

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstNumber := traverseAndGetNumber(l1)
	secondNumber := traverseAndGetNumber(l2)

	result := firstNumber + secondNumber

	splitted := strings.Split(fmt.Sprint(result), "")
	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}

	val1, _ := strconv.Atoi(splitted[len(splitted)-1])
	firstThird := ListNode{Val: val1, Next: nil}

	nodeToPoint := &firstThird
	for i := 1; (len(splitted) - i - 1) >= 0; i++ {
		val2, _ := strconv.Atoi(splitted[len(splitted)-i-1])
		nextNode := ListNode{Val: val2, Next: nodeToPoint}
		nodeToPoint = &nextNode
	}

	return nodeToPoint
}

func main() {
	firstThird := ListNode{Val: 3, Next: nil}
	firstSecond := ListNode{Val: 4, Next: &firstThird}
	firstFirst := ListNode{Val: 2, Next: &firstSecond}

	secondThird := ListNode{Val: 4, Next: nil}
	secondSecond := ListNode{Val: 6, Next: &secondThird}
	secondFirst := ListNode{Val: 5, Next: &secondSecond}

	addTwoNumbers(&firstFirst, &secondFirst)
}
