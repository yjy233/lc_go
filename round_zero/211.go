package roundzero

type WordDictionary struct {
	Stop bool

	Next [26]*WordDictionary
}

func Constructor211() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	r := this
	for _, w := range word {

		offset := int(w) - int('a')
		if r.Next[offset] == nil {
			r.Next[offset] = &WordDictionary{}
		}

		r = r.Next[offset]
	}

	r.Stop = true
}

func (this *WordDictionary) Search(word string) bool {

	r := this
	for i, w := range word {
		if w == '.' {
			tmp := ""
			if i < len(word)-1 {
				tmp = word[i+1:]
			}

			for j := 0; j < 26; j++ {
				if r.Next[j] == nil {
					continue
				}

				if r.Next[j].Search(tmp) {
					return true
				}
			}

			return false
		}

		offset := int(w) - int('a')
		if r == nil || r.Next[offset] == nil {
			return false
		}

		r = r.Next[offset]

	}

	if r == nil {
		return false
	}
	return r.Stop
}
