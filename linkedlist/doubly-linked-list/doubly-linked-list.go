package doublylinkedlist

import (
	"fmt"
)

type node struct {
	data int32
	next *node
	prev *node
}

// DoublyLinkedList is the struct for hold one head of DLL
type DoublyLinkedList struct {
	head *node
}

// New creates a empty doubly linked list and returns the reference to it
func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
	}
}

/*
	Insert function creates a node and adds if the head does not exist
	or traverses to the end of the linked list until the last node
	and attaches the new node
*/
func (dll *DoublyLinkedList) Insert(data int32) {
	newNode := &node{
		data: data,
		next: nil,
		prev: nil,
	}
	if dll.head == nil {
		dll.head = newNode
	} else {
		temp := dll.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
		newNode.prev = temp
	}
}

//Print prints the linked list
func (dll *DoublyLinkedList) Print() {
	temp := dll.head
	if temp == nil {
		return
	}
	for temp != nil {
		fmt.Printf("<- %d -> ", temp.data)
		temp = temp.next
	}
	println()
}

// Length returns the length of the linked list
func (dll *DoublyLinkedList) Length() int32 {
	var count int32 = 0
	temp := dll.head
	for temp != nil {
		count += 1
		temp = temp.next
	}
	return count
}

/*
	defer is a keyword which halts the execution until the rest of the body gets executed.
	In the below function, if the position to be deleted, is more than the
	length of the linked list, then we panic ( quit abruptly causing the program to exit)
	Sometimes, we need to handle and recover from panics and shutdown gracefully.
	The defer statement catches the panic and recovers from it by printing the error.
*/
func (dll *DoublyLinkedList) DeleteAtGivenPosition(position int32) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	lengthOfdll := dll.Length()
	if lengthOfdll < position {
		panic("Position out of bound")
	} else {
		if dll.head == nil {
			return
		}

		temp := dll.head
		if position == 1 {
			dll.head = temp.next
			dll.head.prev = nil
			return
		}

		var currentPosition int32 = 1
		for temp != nil {
			if currentPosition+1 == position {
				break
			}
			currentPosition += 1
			temp = temp.next
		}
		if temp.next == nil {
			return
		} else {
			temp.next = temp.next.next
			if temp.next != nil {
				temp.next.prev = temp
			}
		}
	}
}

func (dll *DoublyLinkedList) Reverse() {
	temp := dll.head
	if temp == nil {
		return
	}
	for temp != nil {
		swap := temp.next
		temp.next = temp.prev
		if swap == nil {
			dll.head = temp
		}
		temp.prev = swap
		temp = temp.prev
	}
}
