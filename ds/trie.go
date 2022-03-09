package ds

import "fmt"

type TrieNode struct {
	Kids map[string]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrieNode() *TrieNode {
	var trieNode TrieNode
	trieNode.Kids = make(map[string]*TrieNode)

	return &trieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Search(word string) *TrieNode {
	currKid := t.Root

	for _, char := range word {
		kidNode, exists := currKid.Kids[string(char)]
		if exists {
			currKid = kidNode
		} else {
			return nil
		}
	}

	return currKid
}

func (t *Trie) Insert(word string) *TrieNode {
	currKid := t.Root

	for _, char := range word {
		kidNode, exists := currKid.Kids[string(char)]
		if exists {
			currKid = kidNode
		} else {
			newKid := NewTrieNode()
			currKid.Kids[string(char)] = newKid
			currKid = newKid
		}
	}

	currKid.Kids["*"] = nil

	return currKid
}

func (t *Trie) CollectAllWords(node *TrieNode, word string, words []string) {
	currKid := node
	if currKid == nil {
		currKid = t.Root
	}

	for key, kidNode := range currKid.Kids {
		if key == "*" {
			words = append(words, word)
			fmt.Println(words)
		} else {
			t.CollectAllWords(kidNode, word+key, words)
		}
	}
}
