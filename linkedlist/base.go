package linkedlist

import "fmt"

// BaseImplementer creates an interface for all the base methods a linked list should have
type BaseImplementer interface {
	Insert(data interface{})
	Print()
	DeleteAtGivenPosition(position int32)
	Length()
	Reverse()
}

// HelloLinkedList is a dummy method to see if the package exports correctly
func HelloLinkedList() {
	fmt.Println("Hello Linked List")
}
