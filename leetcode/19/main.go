package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := ListNode{1, nil}
	n2 := ListNode{2, &n1}
	n3 := ListNode{3, &n2}
	n4 := ListNode{4, &n3}
	_ = ListNode{5, &n4}

	result := removeNthFromEnd(&n2, 2)
	fmt.Println(result)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeLength := 0

	counter := head

	for {
		if counter == nil {
			break
		}

		nodeLength++
		counter = counter.Next
	}

	target := head

	for i := 1; i < nodeLength-n; i++ {
		target = target.Next
	}

	if nodeLength == n {
		head = head.Next
	} else if target.Next != nil {
		target.Next = target.Next.Next
	}

	return head
}
