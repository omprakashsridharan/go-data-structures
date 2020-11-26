package main

import (
	"data-structures/linkedlist"
	doublylinkedlist "data-structures/linkedlist/doubly-linked-list"
	singlylinkedlist "data-structures/linkedlist/singly-linked-list"
	"fmt"
)

func performInterfaceMethods(listType string, ll linkedlist.BaseImplementer) {
	fmt.Printf("--- %s ---", listType)
	println()
	ll.Insert(10)
	ll.Print()
	ll.Insert(20)
	ll.Print()
	ll.Insert(30)
	ll.Print()
	ll.Insert(40)
	ll.Print()
	println("Length of singly linked list: ", ll.Length())
	println("Deleteing node at position 4")
	ll.DeleteAtGivenPosition(4)
	ll.Print()
	println("Reversing linked list")
	ll.Reverse()
	ll.Print()
	println("Deleteing node at position 5 which is not existent")
	ll.DeleteAtGivenPosition(5)
	ll.Print()
	println()
}

func main() {
	sll := singlylinkedlist.New()
	performInterfaceMethods("Singly Linked List", sll)
	dll := doublylinkedlist.New()
	performInterfaceMethods("Doubly Linked List", dll)
}
