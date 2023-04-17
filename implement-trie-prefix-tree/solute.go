package solute

// [实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)
type Trie struct {
	children [26]*Trie
	isDone   bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isDone = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		if node == nil {
			return false
		}

		node = node.children[c-'a']
	}

	return node != nil && node.isDone
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		if node == nil {
			return false
		}

		node = node.children[c-'a']
	}

	return node != nil
}
