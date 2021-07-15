package main

import "fmt"

type node struct {
	value    int
	next     *node
	previous *node
}

type linkedList struct {
	head   *node
	tile   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	currentHead := l.head
	n.next = l.head
	if currentHead != nil {
		currentHead.previous = n
	}

	l.head = n
	if l.length == 0 {
		l.tile = n
	}

	l.length++
}

func (l *linkedList) append(n *node) {
	n.previous = l.tile

	if l.tile != nil {
		l.tile.next = n
	}

	l.tile = n
	if l.length == 0 {
		l.head = n
	}

	l.length++
}

func (l linkedList) printValues(direction string) {
	if direction == "front" {
		node := l.head
		for l.length != 0 {
			fmt.Println(node.value)
			node = node.next
			l.length--
		}
	}

	if direction == "rear" {
		node := l.tile
		for l.length != 0 {
			fmt.Println(node.value)
			node = node.previous
			l.length--
		}
	}

}

func main() {
	list := linkedList{}

	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}

	node4 := &node{value: 4}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	list.append(node4)

	fmt.Println(list)
	list.printValues("front")
	list.printValues("rear")
}
