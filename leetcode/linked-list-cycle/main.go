package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	touchedNode := []*ListNode{}
	touchedNode = append(touchedNode, head)

	if head == nil {
		return false
	}

	cursor := head.Next

	for cursor != nil {
		for _, node := range touchedNode {
			if node == cursor {
				return true
			}
		}

		touchedNode = append(touchedNode, cursor)

		cursor = cursor.Next
	}

	return false
}

func main() {}
