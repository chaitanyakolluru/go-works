package main

import "testing"

func TestMergetwoLists(t *testing.T) {

	firstThird := ListNode{Val: 4}
	firstSecond := ListNode{Val: 2, Next: &firstThird}
	firstFirst := ListNode{Val: 1, Next: &firstSecond}

	secondThird := ListNode{Val: 4}
	secondSecond := ListNode{Val: 3, Next: &secondThird}
	secondFirst := ListNode{Val: 1, Next: &secondSecond}

	got := mergeTwoLists(&firstFirst, &secondFirst)

	wantSixth := &ListNode{Val: 4}
	wantFifth := &ListNode{Val: 4, Next: wantSixth}
	wantFourth := &ListNode{Val: 3, Next: wantFifth}
	wantThird := &ListNode{Val: 2, Next: wantFourth}
	wantSecond := &ListNode{Val: 1, Next: wantThird}
	wantFirst := &ListNode{Val: 1, Next: wantSecond}

	tempG := got
	tempW := wantFirst

	for {
		if tempG.Val != tempW.Val {
			t.Errorf("got: %d, want: %d", tempG.Val, tempW.Val)
		}
		tempG = tempG.Next
		tempW = tempW.Next

		if tempG == nil {
			break
		}
	}

}

func TestMergetwoListsEmpty(t *testing.T) {

	firstThird := ListNode{Val: 4}
	firstSecond := ListNode{Val: 2, Next: &firstThird}
	firstFirst := ListNode{Val: 1, Next: &firstSecond}

	secondThird := ListNode{Val: 4}
	secondSecond := ListNode{Val: 3, Next: &secondThird}
	secondFirst := ListNode{Val: 1, Next: &secondSecond}

	got := mergeTwoLists(&firstFirst, &secondFirst)

	wantSixth := &ListNode{Val: 4}
	wantFifth := &ListNode{Val: 4, Next: wantSixth}
	wantFourth := &ListNode{Val: 3, Next: wantFifth}
	wantThird := &ListNode{Val: 2, Next: wantFourth}
	wantSecond := &ListNode{Val: 1, Next: wantThird}
	wantFirst := &ListNode{Val: 1, Next: wantSecond}

	tempG := got
	tempW := wantFirst

	for {
		if tempG.Val != tempW.Val {
			t.Errorf("got: %d, want: %d", tempG.Val, tempW.Val)
		}
		tempG = tempG.Next
		tempW = tempW.Next

		if tempG == nil {
			break
		}
	}

}
