type Trie struct {
	hashMap map[rune]*Trie
	isEnd   bool
}

func Constructor() Trie {
	return Trie{
		hashMap: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	current := this
	this.isEnd = true
	for _, letter := range word {
		if next, ok := current.hashMap[letter]; ok {
			current = next
		} else {
			newTrie := Constructor()
			current.hashMap[letter] = &newTrie
			current = &newTrie
		}
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, letter := range word {
		if next, ok := current.hashMap[letter]; ok {
			current = next
		} else {
			return false
		}
	}
	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, letter := range prefix {
		if next, ok := current.hashMap[letter]; ok {
			current = next
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
