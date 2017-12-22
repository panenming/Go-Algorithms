package main

import "fmt"

type doubleNode struct {
	val  int
	next *doubleNode
	prev *doubleNode
}

type doubbleLinkedList struct {
	head *doubleNode
}

func newDoubleNode(val int) *doubleNode {
	n := &doubleNode{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *doubbleLinkedList) addAtBeg(val int) {
	n := newDoubleNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *doubbleLinkedList) addAtEnd(val int) {
	n := newDoubleNode(val)
	if ll.head == nil {
		ll.head = n
		return
	}
	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
	n.prev = cur
}

func (ll *doubbleLinkedList) delAtBeg() int {
	if ll.head == nil {
		return -1
	}
	cur := ll.head
	ll.head = cur.next
	if ll.head != nil {
		cur.next.prev = nil
	}
	return cur.val

}

func (ll *doubbleLinkedList) delAtEnd() int {
	// 没有数据
	if ll.head == nil {
		return -1
	}

	//只有一个值
	if ll.head.next == nil {
		return ll.delAtBeg()
	}

	//多个值
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}
	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *doubbleLinkedList) count() int {
	var ctr int = 0
	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}
	return ctr
}

func (ll *doubbleLinkedList) reverse() {
	var prev, next *doubleNode
	cur := ll.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}
	ll.head = prev
}
func (ll *doubbleLinkedList) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

func (ll *doubbleLinkedList) displayReverse() {
	if ll.head == nil {
		return
	}
	var cur *doubleNode
	for cur = ll.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}
