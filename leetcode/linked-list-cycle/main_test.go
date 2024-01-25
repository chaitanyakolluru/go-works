package main

import "testing"

func TestHasCycle(t *testing.T) {
	// creating test linked list
	tail := &ListNode{4, nil}
	third := &ListNode{3, tail}
	second := &ListNode{2, third}
	head := &ListNode{1, second}

	tail.Next = second

	want := true
	got := hasCycle(head)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

}
