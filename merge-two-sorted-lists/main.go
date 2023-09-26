package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func processList(ll *ListNode) []int {
	if ll == nil {
		return []int{}
	}
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

	return listOfDigits
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	listA1 := processList(list1)
	listA2 := processList(list2)

	mergedList := append(listA1, listA2...)
	sort.Slice(mergedList, func(i, j int) bool {
		return mergedList[i] < mergedList[j]
	})

	if len(mergedList) == 0 {
		return nil
	}

	nodeToPoint := &ListNode{Val: mergedList[len(mergedList)-1]}

	for i := 1; (len(mergedList) - i - 1) >= 0; i++ {
		nextNode := ListNode{Val: mergedList[len(mergedList)-i-1], Next: nodeToPoint}
		nodeToPoint = &nextNode
	}

	return nodeToPoint
}

func main() {
	firstThird := ListNode{Val: 4}
	firstSecond := ListNode{Val: 2, Next: &firstThird}
	firstFirst := ListNode{Val: 1, Next: &firstSecond}

	secondThird := ListNode{Val: 4}
	secondSecond := ListNode{Val: 3, Next: &secondThird}
	secondFirst := ListNode{Val: 1, Next: &secondSecond}

	result := mergeTwoLists(&firstFirst, &secondFirst)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}
