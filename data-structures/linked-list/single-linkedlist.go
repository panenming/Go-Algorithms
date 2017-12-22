package main

import "fmt"

type singleNode struct {
	val  int
	next *singleNode
}

type singleLinkedList struct {
	head *singleNode
}

func newSingleNode(val int) *singleNode {
	n := &singleNode{}
	n.next = nil
	n.val = val
	return n
}

func (ll *singleLinkedList) addAtBeg(val int) {
	n := newSingleNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *singleLinkedList) addAtEnd(val int) {
	n := newSingleNode(val)
	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
}

func (ll *singleLinkedList) delAtBeg() int {
	if ll.head == nil {
		return -1
	}
	cur := ll.head
	ll.head = cur.next
	return cur.val
}

func (ll *singleLinkedList) delAtEnd() int {
	if ll.head == nil {
		return -1
	}
	if ll.head.next == nil {
		return ll.delAtBeg()
	}

	// 多个值
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}
	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *singleLinkedList) count() int {
	var ctr int = 0
	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}
	return ctr

}

func (ll *singleLinkedList) reverse() {
	var prev, next *singleNode
	cur := ll.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	ll.head = prev
}

func (ll *singleLinkedList) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
	fmt.Print("\n")

}
