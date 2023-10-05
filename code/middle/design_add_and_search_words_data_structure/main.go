package design_add_and_search_words_data_structure

type Node struct {
	IsWord   bool
	Children map[rune]*Node
}

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{Children: make(map[rune]*Node)}}
}

func (w *WordDictionary) AddWord(word string) {
	curr := w.Root
	for _, c := range word {
		if _, ok := curr.Children[c]; !ok {
			curr.Children[c] = &Node{Children: make(map[rune]*Node)}
		}
		curr = curr.Children[c]
	}
	curr.IsWord = true
}

func (w *WordDictionary) Search(word string) bool {
	return w.search(word, w.Root)
}

func (w *WordDictionary) search(word string, node *Node) bool {
	if len(word) == 0 {
		return node.IsWord
	}
	if word[0] == '.' {
		for _, child := range node.Children {
			if w.search(word[1:], child) {
				return true
			}
		}
		return false
	}
	if _, ok := node.Children[rune(word[0])]; !ok {
		return false
	}
	return w.search(word[1:], node.Children[rune(word[0])])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
