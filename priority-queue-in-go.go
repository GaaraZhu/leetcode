package leetcode

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name     string
	Priority int

	index int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	length := len(*pq)
	item := x.(Item)
	item.index = length
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	tmp := pq[i]
	pq[i] = pq[j]
	pq[j] = tmp
	pq[i].index = j
	pq[j].index = i
}

func main() {
	listItems := []Item{
		{Name: "Carrot", Priority: 30},
		{Name: "Spinach", Priority: 5},
		{Name: "Potato", Priority: 45},
		{Name: "Rice", Priority: 100},
	}

	priorityQueue := make(PriorityQueue, len(listItems))
	for i, item := range listItems {
		priorityQueue[i] = item
		priorityQueue[i].index = i
	}

	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(Item)
		fmt.Println("Name: %s, Priority: %v", item.Name, item.Priority)
	}
}
