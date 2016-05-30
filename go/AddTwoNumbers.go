package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(a int, b int, carry *int) int {
	c := a + b + carry
	carry = c / 10
	return c % 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil && l2 == nil:
		return l1
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var (
		head  = l1
		pre   = l1
		carry int
	)

	pre.Val = add(l1.Val, l2.Val, &carry)
	head = pre
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil && l2 != nil {
		pre.Next.Val = add(l1.Val, l2.Val, &carry)
		pre = pre.Next
		l1 = l1.next
		l2 = l2.Next
	}

	if l2 != nil {
		pre.Next = l2
	}

	for carry != 0 {
		if pre.Next != nil {
			pre.Next.Val = add(pre.Next.Val, 0, &carry)
			pre = pre.Next
		} else {
			pre.Next = &ListNode{carry, nil}
			break
		}
	}

	return head
}
