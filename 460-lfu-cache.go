package main

import (
	"fmt"
	"math"
)

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}

type node struct {
	key, value      int
	frequency       uint16
	prev, next, end *node
}

func (n *node) reset() {
	n.next = nil
	n.prev = nil
	n.end = nil
}

type LFUCache struct {
	capacity     int
	cache        map[int]*node
	frequencyMap map[uint16]*node
	frequencies  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, map[int]*node{}, map[uint16]*node{}, 0}
}

func (this *LFUCache) Get(key int) int {
	defer this.State()

	n, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 1. remove from the old frequencyMap element
	this.RemoveFromFrequencyMap(n)

	// 2. increment the frequency
	n.frequency = n.frequency + 1

	// 3. add to the new frequencyMap element
	this.AddToFrequencyMap(n)

	return n.value
}

func (this *LFUCache) Put(key int, value int) {
	defer this.State()

	if _, ok := this.cache[key]; !ok {
		if len(this.cache) >= this.capacity {
			this.Evict()
		}
		n := &node{key, value, 0, nil, nil, nil}
		this.cache[key] = n
		this.AddToFrequencyMap(n)
	}
}

func (this *LFUCache) Evict() {
	var smallestFreq float64
	smallestFreq = float64(this.frequencies & (-this.frequencies))
	if smallestFreq != 0 {
		smallestFreq = math.Log2(smallestFreq)
	}
	head, ok := this.frequencyMap[uint16(smallestFreq)]
	if !ok {
		fmt.Println("***************")
		fmt.Println(smallestFreq)
		this.State()
		panic("something must be wrong again!!!")
	}

	end := head.end
	delete(this.cache, end.key)
	this.RemoveFromFrequencyMap(head.end)
}

func (this *LFUCache) AddToFrequencyMap(n *node) {
	oldHeadN, ok := this.frequencyMap[n.frequency]
	if !ok {
		n.end = n
		this.frequencyMap[n.frequency] = n
		this.frequencies = this.frequencies | 1<<n.frequency
		return
	}
	oldHeadN.prev = n
	n.next = oldHeadN
	n.end = oldHeadN.end
	oldHeadN.end = nil
	this.frequencyMap[n.frequency] = n
}

func (this *LFUCache) RemoveFromFrequencyMap(n *node) {
	if n.prev == nil {
		if n.next == nil {
			delete(this.frequencyMap, n.frequency)
			this.frequencies = this.frequencies ^ 1<<n.frequency
			return
		}
		n.next.prev = nil
		n.next.end = n.end
		this.frequencyMap[n.frequency] = n.next
	} else if n.next == nil {
		headN, ok := this.frequencyMap[n.frequency]
		if !ok {
			panic("something must be wrong!!!")
		}
		n.prev.next = nil
		headN.end = n.prev
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	n.reset()
}

func (this *LFUCache) State() {
	fmt.Println("-------------LFUCache Internal State------------")
	fmt.Println(fmt.Sprintf("LFUCache capacity: %v", this.capacity))
	fmt.Println("LFUCache cache:")
	for k, n := range this.cache {
		fmt.Println(fmt.Sprintf("      key: %v, value: %v, frequency: %v", k, n.value, n.frequency))
	}

	fmt.Println("LFUCache frequencyMap:")
	for k, v := range this.frequencyMap {
		fmt.Println(fmt.Sprintf("      key: %v, head: %v", k, v.value))
	}

	fmt.Println(fmt.Sprintf("LFUCache frequencies: %v", this.frequencies))
	fmt.Println("------------------------------------------------")
}
