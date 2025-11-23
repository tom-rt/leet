package main

import (
	"fmt"
	"testing"
)

func TestMergeSortedList(t *testing.T) {
	l11 := ListNode{
		Val:  1,
		Next: nil,
	}
	l12 := ListNode{
		Val:  2,
		Next: nil,
	}
	l13 := ListNode{
		Val:  4,
		Next: nil,
	}
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = nil

	l21 := ListNode{
		Val:  1,
		Next: nil,
	}
	l22 := ListNode{
		Val:  3,
		Next: nil,
	}
	l23 := ListNode{
		Val:  4,
		Next: nil,
	}
	l21.Next = &l22
	l22.Next = &l23
	l23.Next = nil

	ret := mergeTwoLists(&l11, &l21)
	printList(ret)
}

func printList(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
