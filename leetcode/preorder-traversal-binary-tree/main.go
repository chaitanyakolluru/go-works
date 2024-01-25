package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	if nil == root {
		return result
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		x := q[len(q)-1]
		q = q[0 : len(q)-1]

		result = append(result, x.Val)
		if x.Right != nil {
			q = append(q, x.Right)
		}
		if x.Left != nil {
			q = append(q, x.Left)
		}
	}

	return result
}
