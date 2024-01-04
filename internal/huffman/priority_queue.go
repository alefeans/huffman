package huffman

import (
	"container/heap"
)

type Weighter interface {
	Weight() int
}

type PriorityQueue []Weighter

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// The lowest weight has the highest priority (min heap)
	return pq[i].Weight() < pq[j].Weight()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item any) {
	*pq = append(*pq, item.(Weighter))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// NewPriorityQueue returns a PriorityQueue of Nodes with min heap property
// based on a rune frequency map.
func NewPriorityQueue(runeFrequency map[rune]int) *PriorityQueue {
	i := 0
	pq := make(PriorityQueue, len(runeFrequency))

	for r, frequency := range runeFrequency {
		pq[i] = &Node{
			value:  r,
			weight: frequency,
		}
		i++
	}

	heap.Init(&pq)
	return &pq
}
