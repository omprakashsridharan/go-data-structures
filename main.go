package main

import (
	singlylinkedlist "data-structures/linkedlist/singly-linked-list"
)

func main() {
	sll := singlylinkedlist.New()
	sll.Insert(10)
	sll.Print()
	sll.Insert(20)
	sll.Print()
	sll.Insert(30)
	sll.Print()
	sll.DeleteAtGivenPosition(3)
	sll.Print()
	println("Length of singly linked list: ", sll.Length())
	println("Deleteing node at position 3 which is not existent")
	sll.DeleteAtGivenPosition(3)
	println("Reversing linked list")
	sll.Reverse()
	sll.Print()
	println("Program exited successfully")
}
