package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	curr1 := list1
	curr2 := list2
	if curr1.Val < curr2.Val {
		head = curr1
		curr1 = curr1.Next
		head.Next = nil
	} else {
		head = curr2
		curr2 = curr2.Next
		head.Next = nil
	}

	for curr1 != nil || curr2 != nil {
		if curr1 == nil {
			mergeList(head, curr2)
			return head
		} else if curr2 == nil {
			mergeList(head, curr1)
			return head
		}

		if curr1.Val < curr2.Val {
			newNode := curr1
			curr1 = curr1.Next
			appendNodeToList(head, newNode)
		} else {
			newNode := curr2
			curr2 = curr2.Next
			appendNodeToList(head, newNode)
		}
	}
	return head
}

func appendNodeToList(head *ListNode, newNode *ListNode) {
	idx := head
	for idx.Next != nil {
		idx = idx.Next
	}
	idx.Next = newNode
	newNode.Next = nil
}

func mergeList(l1 *ListNode, l2 *ListNode) {
	idx := l1
	for idx.Next != nil {
		idx = idx.Next
	}
	idx.Next = l2
}
