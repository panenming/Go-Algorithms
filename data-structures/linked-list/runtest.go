package main

import (
	"fmt"
)

func main() {
	ll := doubbleLinkedList{}

	ll.addAtBeg(10)
	ll.addAtEnd(20)
	ll.display()
	ll.addAtBeg(30)
	ll.display()

	ll.reverse()
	ll.display()
	ll.displayReverse()

	fmt.Print(ll.delAtBeg(), "\n")
	fmt.Print(ll.delAtEnd(), "\n")
	fmt.Print("Display")
	ll.display()
	fmt.Print(ll.delAtBeg(), "\n")
	ll.display()
	fmt.Print(ll.delAtBeg(), "\n")
	ll.display()

	fmt.Println("singleLinkedList-----")
	sin := singleLinkedList{}
	sin.addAtBeg(10)
	sin.addAtEnd(20)
	sin.display()
	sin.addAtBeg(30)
	sin.display()
	sin.reverse()
	sin.display()
	fmt.Print(sin.delAtBeg(), "\n")
	sin.display()
	fmt.Print(sin.delAtEnd(), "\n")
	sin.display()
}
