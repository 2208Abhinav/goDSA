package main

import "fmt"

func main() {
	var nodeValue string
	head := Node{}
	fmt.Println("Enter the elements of singly linked list")
	fmt.Println("(Type quit and then enter to exit)")
	// Linked list builder logic
	for true {
		fmt.Scan(&nodeValue)
		if nodeValue == "quit" {
			break
		}
		head.insertNodeAtTheEnd(nodeValue)
	}
	printLinkedList(head)
	fmt.Println(getListLength(&head))
}
