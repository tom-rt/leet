package main

import (
	"fmt"
	"testing"
)

func binaryTreePathsTest(t *testing.T) {
	n1 := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	ret := binaryTreePaths(&n1)
	fmt.Println(ret)
}
