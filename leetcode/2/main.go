package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := ListNode{3, nil}
	n2 := ListNode{4, &n1}
	n3 := ListNode{2, &n2}

	n4 := ListNode{4, nil}
	n5 := ListNode{6, &n4}
	n6 := ListNode{5, &n5}

	result := addTwoNumbers(&n3, &n6)
	fmt.Println(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	carry := 0
	num := 0

	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		num += carry

		cur.Next = &ListNode{num % 10, nil}
		cur = cur.Next

		if num >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		num = 0
	}

	return head.Next
}
