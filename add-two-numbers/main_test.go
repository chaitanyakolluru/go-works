package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	firstThird := ListNode{Val: 3, Next: nil}
	firstSecond := ListNode{Val: 4, Next: &firstThird}
	firstFirst := ListNode{Val: 2, Next: &firstSecond}

	secondThird := ListNode{Val: 4, Next: nil}
	secondSecond := ListNode{Val: 6, Next: &secondThird}
	secondFirst := ListNode{Val: 5, Next: &secondSecond}

	got := addTwoNumbers(&firstFirst, &secondFirst)

	wantThird := ListNode{Val: 8, Next: nil}
	wantSecond := ListNode{Val: 0, Next: &wantThird}
	wantFirst := ListNode{Val: 7, Next: &wantSecond}

	if got.Val != wantFirst.Val {
		t.Errorf("got: %d, want: %d", got.Val, wantFirst.Val)
	}

	gotSecond := got.Next
	if gotSecond.Val != wantSecond.Val {
		t.Errorf("got: %d, want: %d", gotSecond.Val, wantSecond.Val)
	}

	gotThird := gotSecond.Next
	if gotThird.Val != wantThird.Val {
		t.Errorf("got: %d, want: %d", gotThird.Val, wantThird.Val)
	}

}
