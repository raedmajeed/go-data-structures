package trie

import "fmt"

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

type trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
	}
}

func (t *trie) Insert(word string) {
	current := t.root
	for _, r := range word {
		if _, ok := current.children[r]; !ok {
			current.children[r] = NewNode()
		}
		fmt.Println(current.children[r], r)
		current = current.children[r]
	}
	current.isEnd = true
}

func (t *trie) Check(word string) bool {
	current := t.root
	for _, r := range word {
		if _, ok := current.children[r]; !ok {
			return false
		}
		current = current.children[r]
	}
	return current.isEnd
}

func (t *trie) Delete(word string) {
	DeleteRec(word, 0, t.root)
}

func DeleteRec(word string, length int, curr *Node) bool {
	if len(word) == length {
		if curr.isEnd {
			curr.isEnd = false
			return len(curr.children) == 0
		}
		return false
	}

	w := rune(word[length])
	check := DeleteRec(word, length, curr.children[w])
	if !check {
		return false
	}
	delete(curr.children, w)
	return len(curr.children) == 0 && !curr.isEnd
}

func TrieOperations() {
	t := &trie{
		root: NewNode(),
	}
	t.Insert("raed")
	fmt.Println(t.Check("raed"))
}
