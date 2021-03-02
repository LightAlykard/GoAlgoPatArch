package main

import (
	"fmt"
	"log"
)

type Node struct {
	next *Node
	prev *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Print() {
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		fmt.Printf("Value %d, Prev: %v, Next: %v \n", tmp.Data, tmp.prev, tmp.next)
	}

}

func (l *List) Len() {
	fmt.Printf("Length %d \n", l.len)
}

func (l *List) Find(elem int) *Node {
	if l.head != nil {
		for tmp := l.head; tmp != nil; tmp = tmp.next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	fmt.Println("Element not find")
	return nil
}

func (l *List) PushBack(node *Node) {
	prev := l.tail
	l.len++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	prev.next = node
	node.prev = prev
	l.tail = node
}

func (l *List) PushFront(node *Node) {

	prev := l.head

	l.len++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = prev
	prev.prev = node
	l.head = node

}

func (l *List) PopBack() {

	if l.len <= 0 {
		log.Fatal("PopBack() called on empty queue")
	}

	del := l.tail
	node := del.prev
	node.next = nil
	del.prev = nil
	l.tail = node
	l.len--
}

func (l *List) PopFront() {
	if l.len <= 0 {
		log.Fatal("PopFront() called on empty queue")
	}

	del := l.head
	head := l.head.next

	head.prev = nil
	del.next = nil
	l.head = head
	l.len--
}

func main() {
	l := &List{}

	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}
	node4 := &Node{Data: 4}

	fmt.Println("PushBack")
	l.PushBack(node1)
	l.PushBack(node2)
	l.PushBack(node3)

	l.Print()

	fmt.Println("\nPushFront")
	l.PushFront(node4)

	l.Print()

	findChislo := l.Find(1)
	if findChislo != nil {
		fmt.Printf("Element find: Value %d, Prev %v, Next %v \n", findChislo.Data, findChislo.prev.Data, findChislo.next.Data)
	}

	fmt.Println("\nPopBack")

	l.PopBack()
	l.Print()

	fmt.Println("\nPopFront")

	l.PopFront()
	l.Print()

}
