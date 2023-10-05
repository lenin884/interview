package implement_trie_prefix

type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.Root

	for _, c := range word {
		if _, ok := curr.Children[c]; !ok {
			curr.Children[c] = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.Children[c]
	}
	curr.IsWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t.Root

	for _, c := range word {
		if _, ok := curr.Children[c]; !ok {
			return false
		}
		curr = curr.Children[c]
	}

	return curr.IsWord
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.Root

	for _, c := range prefix {
		if _, ok := curr.Children[c]; !ok {
			return false
		}
		curr = curr.Children[c]
	}

	return true
}
