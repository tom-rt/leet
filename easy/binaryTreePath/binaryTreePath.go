package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	fmt.Println(root)
	if root.Left == nil && root.Right == nil {
		// This is a leaf

	} else {

	}
	return []string{}
}
