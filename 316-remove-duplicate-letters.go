package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetters(""))         //""
	fmt.Println(removeDuplicateLetters("abc"))      //abc
	fmt.Println(removeDuplicateLetters("adbdc"))    //abdc
	fmt.Println(removeDuplicateLetters("aaaaa"))    //a
	fmt.Println(removeDuplicateLetters("bcabc"))    //abc
	fmt.Println(removeDuplicateLetters("cabccab"))  //abc
	fmt.Println(removeDuplicateLetters("cbacdcbc")) //acdb
}

func removeDuplicateLetters(s string) string {
	lc := map[byte]int{}

	for i := 0; i < len(s); i++ {
		c, ok := lc[s[i]]
		if !ok {
			lc[s[i]] = 1
			continue
		}
		lc[s[i]] = c + 1
	}

	letters := stack{}
	existsInStack := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		cc, ok := lc[s[i]]
		if !ok {
			panic("something is wrong!")
		}

		_, ok = existsInStack[s[i]]
		if ok {
			lc[s[i]] = cc - 1
			continue
		}

		for letters.len() != 0 {
			top := letters.peek()
			c, ok := lc[top]
			if !ok {
				panic("something is wrong!!!")
			}
			if top >= s[i] && c > 0 {
				letters.pop()
				delete(existsInStack, top)
			} else {
				break
			}
		}

		letters.push(s[i])
		existsInStack[s[i]] = struct{}{}
		lc[s[i]] = cc - 1
	}

	var res string
	for letters.len() != 0 {
		top := letters.pop()
		res = string(top) + res
	}

	return res
}

type stack struct {
	elements []byte
}

func (s *stack) push(b byte) {
	s.elements = append(s.elements, b)
}

func (s *stack) len() int {
	return len(s.elements)
}

func (s *stack) peek() byte {
	l := s.len()
	if l == 0 {
		panic("trying to peek on empty stack")
	}
	return s.elements[l-1]
}

func (s *stack) pop() byte {
	l := s.len()
	if l == 0 {
		panic("trying to pop on empty stack")
	}

	returnValue := s.elements[l-1]
	var newEle []byte
	for i := 0; i < l-1; i++ {
		newEle = append(newEle, s.elements[i])
	}
	s.elements = newEle

	return returnValue
}
