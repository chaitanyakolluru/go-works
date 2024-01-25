package main

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	leftLeftRoot := &TreeNode{Val: 4}
	rightLeftRoot := &TreeNode{Val: 5}
	leftRoot := &TreeNode{Val: 2, Left: leftLeftRoot, Right: rightLeftRoot}

	leftRightRoot := &TreeNode{Val: 6}
	rightRightRoot := &TreeNode{Val: 7}
	rightRoot := &TreeNode{Val: 3, Left: leftRightRoot, Right: rightRightRoot}
	root := &TreeNode{Val: 1, Left: leftRoot, Right: rightRoot}

	want := []int{1, 2, 4, 5, 3, 6, 7}
	got := preOrderTraversal(root)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

}
