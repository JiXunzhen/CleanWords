// Package Trie provides ...
package Trie

import "fmt"

type Trie struct {
	/**
	 * 需要其他的数据再说
	 */
	next   map[rune]*Trie
	fail   *Trie
	isWord bool
}

type NodeList struct {
	next *NodeList
	prev *NodeList

	val *Trie
}

func nodPush(tail *NodeList, trie *Trie) *NodeList {
	if tail != nil {
		nod := new(NodeList)
		nod.val = trie

		tail.next = nod
		nod.prev = tail
		return nod
	}
	return nil
}

func nodPop(root *NodeList) *NodeList {
	if root != nil && root.next != nil {
		root.next.prev = nil
		return root.next
	}
	return nil
}
func NewTrie() *Trie {
	trie := new(Trie)
	trie.next = make(map[rune]*Trie)
	return trie
}

func (root *Trie) Insert(key []rune) {
	idx := 0
	length := len(key)

	p := root
	for idx < length {
		if _, in := p.next[key[idx]]; in {
			p = p.next[key[idx]]
		} else {
			nod := NewTrie()
			p.next[key[idx]] = nod
			p = nod
		}
		idx++
	}

	if p != root {
		p.isWord = true
	}
}

func (root *Trie) BuildFailPointer() {
	var pre *Trie

	list := new(NodeList)
	list.val = root
	tail := list

	for list != nil {
		nod := list.val
		for key, val := range nod.next {
			if nod == root {
				val.fail = root
			} else {
				pre = nod.fail

				for pre != nil {
					if nnod, in := pre.next[key]; in {
						val.fail = nnod
						break
					}
					pre = pre.fail
				}
				if pre == nil {
					val.fail = root
				}
			}
			tail = nodPush(tail, val)
		}
		list = nodPop(list)
	}
}

func (root *Trie) IsExist(key []rune) bool {
	idx := 0
	length := len(key)
	p := root

	for idx < length {
		k := key[idx]

		for {
			_, in := p.next[k]
			if !in && p != root {
				p = p.fail
			} else {
				p = p.next[k]
				break
			}
		}
		if p == nil {
			p = root
		}
		if p.isWord {
			return true
		}

		idx++
	}
	return false
}

func (root *Trie) ToString() {
	root.formatPrint(1, 'r')
}
func (nod *Trie) formatPrint(deep int, key rune) {
	for i := 0; i < deep; i++ {
		fmt.Print("-")
	}
	fmt.Println(string(key), nod.isWord)
	for key, val := range nod.next {
		val.formatPrint(deep+1, key)
	}
}
