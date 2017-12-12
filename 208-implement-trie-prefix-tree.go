package main

import "fmt"

func main() {
	//true,true,true,false
	t := Constructor1()
	t.Insert("abc")
	fmt.Println(t.Search("abc"))
	fmt.Println(t.StartsWith("ab"))
	fmt.Println(t.StartsWith("abc"))
	fmt.Println(t.StartsWith("bc"))

	//false,true
	t2 := Constructor1()
	t2.Insert("ab")
	fmt.Println(t2.Search("a"))
	fmt.Println(t2.StartsWith("a"))

	//true,true
	t3 := Constructor1()
	t3.Insert("a")
	fmt.Println(t3.Search("a"))
	fmt.Println(t3.StartsWith("a"))

	//false,true,true,true
	t4 := Constructor1()
	t4.Insert("ab")
	fmt.Println(t4.Search("a"))
	fmt.Println(t4.Search("ab"))
	fmt.Println(t4.StartsWith("a"))
	fmt.Println(t4.StartsWith("ab"))

	//true,false,true,true
	t5 := Constructor1()
	t5.Insert("abc")
	fmt.Println(t5.Search("abc"))
	fmt.Println(t5.Search("ab"))
	t5.Insert("ab")
	fmt.Println(t5.Search("ab"))
	t5.Insert("ab")
	fmt.Println(t5.Search("ab"))

	//false,true,false
	t6 := Constructor1()
	t6.Insert("app")
	t6.Insert("apple")
	t6.Insert("beer")
	t6.Insert("add")
	t6.Insert("jam")
	t6.Insert("rental")
	fmt.Println(t6.Search("apps"))
	fmt.Println(t6.Search("app"))
	fmt.Println(t6.Search("ad"))
}

type Trie struct {
	leaves      map[byte]*Trie
	isEndOfWord bool
}

/** Initialize your data structure here. */
func Constructor1() Trie {
	return Trie{map[byte]*Trie{}, false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.isEndOfWord = true
		return
	}

	t, ok := this.leaves[word[0]]
	if !ok {
		t = &Trie{map[byte]*Trie{}, false}
		this.leaves[word[0]] = t
	}
	t.Insert(word[1:len(word)])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.isEndOfWord
	}
	l, ok := this.leaves[word[0]]
	if !ok {
		return false
	}
	return l.Search(word[1:len(word)])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	l, ok := this.leaves[prefix[0]]
	if !ok {
		return false
	}
	return l.StartsWith(prefix[1:len(prefix)])
}
