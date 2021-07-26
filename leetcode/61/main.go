package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := ListNode{4, nil}
	n2 := ListNode{3, &n1}
	n3 := ListNode{2, &n2}
	n4 := ListNode{1, &n3}

	result := rotateRight(&n4, 10)
	printVals(result)
}

func printVals(head *ListNode) {
	for {
		fmt.Println(head.Val)
		head = head.Next
		if head == nil {
			break
		}
	}

}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	counter := head

	for {
		length++
		if counter.Next == nil {
			k = k % length
			if k != 0 {
				counter.Next = head
			}
			break
		}
		counter = counter.Next
	}

	if k != 0 {
		end := head

		for i := 1; i <= (length - k - 1); i++ {
			end = end.Next
		}
		head = end.Next
		end.Next = nil
	}

	return head

}
