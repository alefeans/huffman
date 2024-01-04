package huffman

import (
	"container/heap"
	"testing"
)

var stubCharsFrequency = map[rune]int{
	1: 32,
	2: 42,
	3: 120,
	4: 7,
	5: 43,
	6: 24,
	7: 37,
	8: 2,
}

func TestPriorityQueue(t *testing.T) {
	want := []struct {
		value  rune
		weight int
	}{
		{value: 8, weight: 2},
		{value: 4, weight: 7},
		{value: 6, weight: 24},
		{value: 1, weight: 32},
		{value: 7, weight: 37},
		{value: 2, weight: 42},
		{value: 5, weight: 43},
		{value: 3, weight: 120}}

	pq := NewPriorityQueue(stubCharsFrequency)

	if pq.Len() != len(want) {
		t.Errorf("got length %d, want %d", pq.Len(), len(want))
	}

	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Node)
		if item.value != want[i].value {
			t.Errorf("got value %d, want %d", item.value, want[i].value)
		}
		if item.weight != want[i].weight {
			t.Errorf("got weight %d, want %d", item.weight, want[i].weight)
		}
		i++
	}
}
