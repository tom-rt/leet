package main

import "testing"

func TestAddTwoNumbersDummy(t *testing.T) {
	// Add two numbers
	node3 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 3, Next: node2}
	node6 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 6, Next: node6}
	node4 := &ListNode{Val: 4, Next: node5}
	sum := addTwoNumbers(node1, node4)
	if sum.Val != 7 {
		t.Fatalf(`Test failed wanted 7, got %d`, sum.Val)
		t.Failed()
	}
}
