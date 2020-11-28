package main

import (
	"data-structures/linkedlist"
	circularlinkedlist "data-structures/linkedlist/circular-linked-list"
	doublylinkedlist "data-structures/linkedlist/doubly-linked-list"
	singlylinkedlist "data-structures/linkedlist/singly-linked-list"
	"data-structures/stack"
	"fmt"
)

func performLinkedListInterfaceMethods(listType string, ll linkedlist.BaseImplementer) {
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
	println()
}

func printStackMethods(s *stack.Stack) {
	println("Height of stack: ", s.Height())
	println("Return -1 on peek if nothing is there in stack")
	println("Peeking now. Result: ", s.Peek())
	println("Pushing 10 to the stack")
	s.Push(10)
	println("Peeking now. Result: ", s.Peek())
	println("Height of stack: ", s.Height())
	println("Pushing 20 to the stack")
	s.Push(20)
	println("Peeking now. Result: ", s.Peek())
	s.Print()
	println("Popping top of the stack")
	if poppedItem, ok := s.Pop(); ok {
		println("Popped Item: ", poppedItem)
	}
	println("Popping top of the stack")
	if poppedItem, ok := s.Pop(); ok {
		println("Popped Item: ", poppedItem)
	}
	println("Popping top of the stack")
	if poppedItem, ok := s.Pop(); ok {
		println("Popped Item: ", poppedItem)
	} else {
		println("Oops! Stack is empty, Nothing to pop")
	}
	println()
}

func main() {
	sll := singlylinkedlist.New()
	performLinkedListInterfaceMethods("Singly Linked List", sll)
	dll := doublylinkedlist.New()
	performLinkedListInterfaceMethods("Doubly Linked List", dll)
	cll := circularlinkedlist.New()
	performLinkedListInterfaceMethods("Circular Linked List", cll)
	s := stack.New()
	printStackMethods(s)
}
