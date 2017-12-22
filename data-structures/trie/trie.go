package main

type trieNode struct {
	children map[rune]*trieNode
	isLeaf   bool
}

func newNode() *trieNode {
	n := &trieNode{}
	n.children = make(map[rune]*trieNode)
	n.isLeaf = false
	return n
}

// 插入一个新的单词到单词树
func (n *trieNode) insert(s string) {
	cur := n
	for _, c := range s {
		next, ok := cur.children[c]
		if !ok {
			next = newNode()
			cur.children[c] = next
		}
		cur = next
	}
	cur.isLeaf = true
}

func (n *trieNode) find(s string) bool {
	cur := n
	for _, c := range s {
		next, ok := cur.children[c]
		if !ok {
			return false
		}
		cur = next
	}
	return true
}
