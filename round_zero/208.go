package roundzero

type Trie struct {
	Over bool
	Next [26]*Trie
}

func Constructor208() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	tmp := this

	for _, c := range word {
		offset := int(c) - int('a')

		if tmp.Next[offset] == nil {
			tmp.Next[offset] = &Trie{}
		}

		tmp = tmp.Next[offset]
	}
	tmp.Over = true
}

func (this *Trie) Search(word string) bool {

	tmp := this

	for _, c := range word {
		offset := int(c) - int('a')

		if tmp.Next[offset] == nil {
			return false
		}

		tmp = tmp.Next[offset]
	}

	return tmp != nil && tmp.Over
}

func (this *Trie) StartsWith(prefix string) bool {

	tmp := this

	for _, c := range prefix {
		offset := int(c) - int('a')

		if tmp.Next[offset] == nil {
			return false
		}

		tmp = tmp.Next[offset]
	}

	return tmp != nil
}
