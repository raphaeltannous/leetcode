type WordDictionary struct {
	hashMap map[rune]*WordDictionary
	isEnd   bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		hashMap: make(map[rune]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, letter := range word {
		if next, ok := current.hashMap[letter]; ok {
			current = next
		} else {
			newWordDict := Constructor()
			current.hashMap[letter] = &newWordDict
			current = current.hashMap[letter]
		}
	}
	current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	current := this
	for x, letter := range word {
		if next, ok := current.hashMap[letter]; ok {
			current = next
		} else if letter == '.' {
			for l, _ := range current.hashMap {
				word = word[:x] + string(l) + word[x+1:]
				if this.Search(word) {
					return true
				}
				word = word[:x] + "." + word[x+1:]
			}
			return false
		} else {
			return false
		}
	}
	return current.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
