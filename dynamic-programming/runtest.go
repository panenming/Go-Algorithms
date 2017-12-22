package main

import "fmt"

func main() {
	word := "aaaabbbba"
	fmt.Printf("%d\n", lpsRec(word, 0, len(word)-1))
	fmt.Printf("%d\n", lpsDp(word))
	fmt.Printf("%d\n", lpsDp2(word))
}
