package singlylinkedlist

import (
	"fmt"
)

type node struct {
	data int32
	next *node
}

// SinglyLinkedList is the struct for hold one head of SLL
type SinglyLinkedList struct {
	head *node
}

// New creates a empty singly linked list and returns the reference to it
func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

/*
	Insert function creates a node and adds if the head does not exist
	or traverses to the end of the linked list until the last node
	and attaches the new node
*/
func (sll *SinglyLinkedList) Insert(data int32) {
	newNode := &node{
		data: data,
		next: nil,
	}
	if sll.head == nil {
		sll.head = newNode
	} else {
		temp := sll.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

//Print prints the linked list
func (sll *SinglyLinkedList) Print() {
	temp := sll.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	println("NULL")
}

// Length returns the length of the linked list
func (sll *SinglyLinkedList) Length() int32 {
	var count int32 = 0
	temp := sll.head
	for temp != nil {
		count += 1
		temp = temp.next
	}
	return count
}

func (sll *SinglyLinkedList) DeleteAtGivenPosition(position int32) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	lengthOfSll := sll.Length()
	if lengthOfSll < position {
		panic("Position out of bound")
	} else {
		if sll.head == nil {
			return
		}

		temp := sll.head
		if position == 1 {
			sll.head = temp.next
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
		}
	}
}

func (sll *SinglyLinkedList) Reverse() {
	current := sll.head
	var previous, next *node = nil, nil
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	sll.head = previous
}
