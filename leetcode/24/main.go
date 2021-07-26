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

	result := swapPairs(&n4)
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	head.Next = swapPairs(head.Next.Next)
	n.Next = head
	return n
}
