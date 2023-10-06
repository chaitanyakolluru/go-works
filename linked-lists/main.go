package main

import "fmt"

type node struct {
	next *node
	data int
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) printListData() {
	toPrint := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.next
	}
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func main() {
	myList := LinkedList{}
	node1 := &node{data: 40}
	node2 := &node{data: 50}
	node3 := &node{data: 60}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	// fmt.Println(myList)
	// myList.printListData()
	myList.deleteWithValue(60)
	myList.printListData()
}
