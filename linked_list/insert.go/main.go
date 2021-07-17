package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (head *Node) append(n int) {
	newNode := Node{n, nil}

	target := head

	for target.next != nil {
		target = target.next
	}

	target.next = &newNode
}

// if m less than Node n's value, insert the m in front of n
func (head *Node) insert(m int) {
	newNode := Node{m, nil}
	pre := head

	if newNode.value < pre.value {
		newNode.value, head.value = head.value, newNode.value
		newNode.next, head.next = head.next, &newNode

		return
	}

	for pre.next != nil {
		if newNode.value < pre.next.value {
			newNode.next = pre.next
			pre.next = &newNode
			break
		} else {
			pre = pre.next
		}
	}
}

func (head Node) printValues() {
	var result []int

	h := head
	for {
		result = append(result, h.value)

		if h.next == nil {
			break
		}

		h = *h.next
	}

	fmt.Println(result)
}

func main() {
	head := Node{1, nil}
	head.append(2)
	head.append(4)

	head.printValues()

	head.insert(3)
	head.printValues()

	head.insert(0)
	head.printValues()

	head.insert(5)
	head.printValues()
}
