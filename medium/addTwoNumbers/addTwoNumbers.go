package main

import (
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var x string
	var y string
	var sum string
	var head *ListNode

	for l1 != nil {
		x = strconv.Itoa(l1.Val) + x
		l1 = l1.Next
	}
	for l2 != nil {
		y = strconv.Itoa(l2.Val) + y
		l2 = l2.Next
	}

	sum = addHugeNumbers(x, y)
	for i := 0; i < len(sum); i++ {
		val, _ := strconv.Atoi(string(sum[i]))
		head = &ListNode{Val: val, Next: head}
	}
	return head
}

func addHugeNumbers(num1, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	num1 = reverse(num1)
	num2 = reverse(num2)

	result := make([]byte, 0, len(num1)+1)
	carry := 0

	for i := 0; i < len(num1); i++ {
		digit1 := int(num1[i] - '0')
		digit2 := 0
		if i < len(num2) {
			digit2 = int(num2[i] - '0')
		}

		sum := digit1 + digit2 + carry
		result = append(result, byte(sum%10)+'0')
		carry = sum / 10
	}

	if carry > 0 {
		result = append(result, byte(carry)+'0')
	}

	return reverse(string(result))
}
