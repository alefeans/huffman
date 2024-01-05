package huffman

import (
	"reflect"
	"testing"
)

func StubHuffmanTree() *Node {
	pq := NewPriorityQueue(map[rune]int{10: 1, 12: 2})
	return NewHuffmanTree(pq)
}

func TestNewHuffmanTree(t *testing.T) {
	got := StubHuffmanTree()

	// Test root node
	if got.value != 0 {
		t.Errorf("got %d, want 0", got.value)
	}
	if got.weight != 3 {
		t.Errorf("got %d, want 3", got.weight)
	}
	if got.code != CodeZero {
		t.Errorf("got %d, want 0", got.code)
	}

	if got.left.value != 10 {
		t.Errorf("got %d, want 0", got.left.value)
	}
	if got.left.weight != 1 {
		t.Errorf("got %d, want 1", got.left.weight)
	}
	if got.left.code != CodeZero {
		t.Errorf("got %d, want 0", got.left.code)
	}

	if got.right.value != 12 {
		t.Errorf("got %d, want 12", got.right.value)
	}
	if got.right.weight != 2 {
		t.Errorf("got %d, want 2", got.right.weight)
	}
	if got.right.code != CodeOne {
		t.Errorf("got %d, want 0", got.right.code)
	}
}

func TestToHeader(t *testing.T) {
	want := "0 10 n n 12 n n"
	got := StubHuffmanTree().ToHeader()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestFromHeader(t *testing.T) {
	input := "0 10 n n 12 n n"
	want := StubHuffmanTree()

	got, err := FromHeader(input)
	if err != nil {
		t.Errorf("got unexpected error with readHeader(%v): %v", input, err)
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToAndFromHeader(t *testing.T) {
	want := StubHuffmanTree().ToHeader()
	tree, _ := FromHeader(want)
	got := tree.ToHeader()

	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestToLookup(t *testing.T) {
	want := map[rune]int{10: 0, 12: 1}
	got := StubHuffmanTree().ToLookup()

	if reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToReverseLookup(t *testing.T) {
	want := map[rune]int{0: 10, 1: 12}
	got := StubHuffmanTree().ToReverseLookup()

	if reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}
