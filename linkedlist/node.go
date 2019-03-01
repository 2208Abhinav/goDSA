package main

import "fmt"

// Node of a singly linked list
type Node struct {
	value string
	next  *Node // Pointer to the next element.
}

func getListLength(head *Node) int {
	length := 0
	for head != nil {
		length++
		head = (*head).next
	}
	return length
}

func (n *Node) setData(value string) {
	(*n).value = value
}

func (n *Node) setNext(next *Node) {
	(*n).next = next
}

// Recursive function to print linked list.
// In the recursive call we are passing the struct
// by value(*(n.next)) and not the pointer because
// by default the function takes value and not pointer.
func printLinkedList(n Node) {
	if n.value == "" {
		fmt.Println(nil)
		return
	}
	fmt.Print(n.value, " --> ")
	if n.next == nil {
		fmt.Println(nil)
		return
	}
	printLinkedList(*(n.next)) // Convert pointer to the value type.
}

// This code block is heavily using pointers and is
// going back and forth between pointers and values
// so get ready to get a heavy dose of pointers.
func (n *Node) insertNodeAtTheEnd(el string) {
	if (*n).value == "" {
		(*n).setData(el)
		return
	}
	if (*n).next == nil {
		nextNode := Node{}
		(*n).next = &nextNode
		(nextNode).insertNodeAtTheEnd(el)
	} else {
		((*n).next).insertNodeAtTheEnd(el)
	}
}
