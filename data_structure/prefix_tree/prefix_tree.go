package prefixtree

// leaf can be string
// array

// char-dict
type Trie struct {
	children map[byte]*Trie
}

func Init(strs []string) *Trie {
	t := &Trie{}
	for i := 0; i < len(strs); i++ {
		t.Insert(strs[i])
	}

	return t
}

func (t *Trie) Insert(s string) {
	root := t
	for i := 0; i < len(s); i++ {
		c := s[i]
		if child, ok := root.children[c]; ok {
			root = child
		} else {
			root.children[c] = &Trie{
				children: make(map[byte]*Trie),
			}

			root = root.children[c]
		}
	}
}

func (t *Trie) Remove(s string) {
	if len(s) == 0 {
		return
	}

	// find
	root := t
	for i := 0; i < len(s)-1; i++ {
		if child, ok := root.children[s[i]]; ok {
			root = child
		} else {
			return
		}
	}

	delete(root.children, s[s[len(s)-1]])
}

func (t *Trie) Search(s string) bool {
	root := t
	for i := 0; i < len(s); i++ {
		if child, ok := root.children[s[i]]; ok {
			root = child
		} else {
			return false
		}
	}

	return true
}
