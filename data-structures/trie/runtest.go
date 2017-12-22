package main

import "fmt"

func main() {
	n := newNode()
	insertWords := [...]string{"nikola", "tesla"}
	for _, w := range insertWords {
		n.insert(w)
	}

	checkWords := map[string]bool{
		"thomas": false,
		"edison": false,
		"nikola": true,
	}
	for k, v := range checkWords {
		ok := n.find(k)
		if ok != v {
			fmt.Println("单词查找树有误！")
		}
		fmt.Printf(
			"%s is %s in the Trie.\n",
			k,
			map[bool]string{true: "", false: "NOT "}[ok])
	}
}
