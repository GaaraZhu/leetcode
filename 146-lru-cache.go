package main

import "fmt"

func main() {
	cache := Constructor(2)
	// test case 1
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1

	// cache.Status()
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4

	// test case 2
	// fmt.Println(cache.Get(2)) // returns -1
	// cache.Put(2, 6)
	// fmt.Println(cache.Get(1)) // returns -1
	// cache.Put(1, 5)
	// // cache.Status()
	// cache.Put(1, 2)
	// // cache.Status()

	// fmt.Println(cache.Get(1)) // returns 2
	// fmt.Println(cache.Get(2)) // returns 6

	// test case 3
	// cache.Put(2, 1)
	// cache.Put(1, 1)

	// cache.Put(2, 3)
	// cache.Status()
	// cache.Put(4, 1)

	// fmt.Println(cache.Get(1)) // returns -1
	// fmt.Println(cache.Get(2)) // returns 3

	// test case 4
	// cache = Constructor(1)
	// cache.Put(2, 1)
	// fmt.Println(cache.Get(2)) // returns 1
	// // cache.Status()
	// cache.Put(3, 2)

	// fmt.Println(cache.Get(2)) // returns -1
	// fmt.Println(cache.Get(3)) // returns 2
}

type node struct {
	key        int
	value      int
	prev, next *node
}

func (n *node) setPrev(p *node) {
	if p == n {
		return
	}
	n.prev = p
	if p != nil {
		p.next = n
	}
}

func (n *node) setNext(ne *node) {
	if ne == n {
		return
	}
	n.next = ne
	if ne != nil {
		ne.prev = n
	}
}

type LRUCache struct {
	capacity int
	values   map[int]*node
	rootNode *node
	endNode  *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*node, capacity), nil, nil}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.values[key]
	if !ok {
		return -1
	}
	this.updateOrder(v)

	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.values[key]
	if !ok {
		if len(this.values) == this.capacity {
			this.Evict()
		}
		n = &node{key, value, nil, nil}
	} else {
		n.value = value
	}
	this.values[key] = n

	this.updateOrder(n)
}

func (this *LRUCache) updateOrder(n *node) {
	if n.prev != nil {
		n.prev.setNext(n.next)
		if this.endNode == n {
			if n.next != nil {
				this.endNode = n.next
			} else if n.prev != nil {
				this.endNode = n.prev
			}
		}
	}
	n.setNext(this.rootNode)
	this.rootNode = n
	n.setPrev(nil)

	if this.endNode == nil {
		this.endNode = n
	}
}

func (this *LRUCache) Evict() {
	delete(this.values, this.endNode.key)
	nextEnd := this.endNode.prev
	if nextEnd != nil {
		nextEnd.next = nil
	}
	this.endNode.prev = nil
	this.endNode = nextEnd
}

func (this *LRUCache) Status() {
	fmt.Println("----map data-----")
	for k, v := range this.values {
		fmt.Println(fmt.Sprintf("key: %v, value: %v", k, v))
	}

	fmt.Println("----root-to-end node order-----")
	tmp := this.rootNode
	current := this.rootNode
	i := 0
	for {
		if (current == nil || current == tmp) && i != 0 {
			break
		}
		fmt.Println(fmt.Sprintf("key: %v", current.key))
		current = current.next
		i++
	}

	fmt.Println("----reverse node order-----")
	tmp = this.endNode
	current = this.endNode
	i = 0
	for {
		if (current == nil || current == tmp) && i != 0 {
			break
		}
		fmt.Println(fmt.Sprintf("key: %v", current.key))
		current = current.prev
		i++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
