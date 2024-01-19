package huffman

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type Code byte

const (
	CodeZero      Code = 0
	CodeOne       Code = 1
	NodeSeparator      = " "
)

type Node struct {
	value  rune
	weight int
	code   Code
	left   *Node
	right  *Node
}

// Implements the Weighter interface required on the PriorityQueue.
func (h *Node) Weight() int {
	return h.weight
}

// NewHuffmanTree builds a HuffmanTree based on the order of weights from the
// priority queue (ref: https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html).
func NewHuffmanTree(pq heap.Interface) *Node {
	var tree *Node

	for pq.Len() > 1 {
		minNode := heap.Pop(pq).(*Node)
		nextMinNode := heap.Pop(pq).(*Node)

		// Assign Huffman Code to edges
		minNode.code = CodeZero
		nextMinNode.code = CodeOne

		tree = &Node{
			weight: minNode.weight + nextMinNode.weight,
			left:   minNode,
			right:  nextMinNode,
		}

		heap.Push(pq, tree)
	}

	return tree
}

// ToHeader serializes the tree in a string format using pre-order traversal.
// The Node's weight is not necessary for either compression and decompression.
func (h *Node) ToHeader() string {
	serialized := []string{}
	serialize(h, &serialized)
	return strings.Join(serialized, NodeSeparator)
}

func serialize(h *Node, path *[]string) {
	if h == nil {
		*path = append(*path, "n")
		return
	}

	*path = append(*path, fmt.Sprintf("%d", h.value))
	serialize(h.left, path)
	serialize(h.right, path)
}

// FromHeader deserializes the header string in a Huffman tree.
func FromHeader(header string) (*Node, error) {
	splitted := strings.Split(header, NodeSeparator)
	return deserialize(&splitted)
}

func deserialize(serialized *[]string) (*Node, error) {
	token := (*serialized)[0]
	*serialized = (*serialized)[1:] // Advance slice to go to the next token

	if token == "n" {
		return nil, nil
	}

	value, err := strconv.Atoi(token)
	if err != nil {
		return nil, err
	}

	curr := &Node{
		value: rune(value),
	}

	left, err := deserialize(serialized)
	if err != nil {
		return nil, err
	}

	if left != nil {
		left.code = CodeZero
	}
	curr.left = left

	right, err := deserialize(serialized)
	if err != nil {
		return nil, err
	}

	if right != nil {
		right.code = CodeOne
	}
	curr.right = right

	return curr, nil
}

// ToLookup assigns the Huffman Code path on the tree to each rune to
// generate a lookup map, ex:
// The tree below (using chars instead of runes to simplify the example):
//
//	   3
//	  0  1
//	 A    8
//           0  1
//	    C    D
//
// Becomes:
// {A: 0, C: 10, B: 11}
func (h *Node) ToLookup() map[rune]string {
	table := make(map[rune]string)
	assignCodePathToChars(h, "", table)
	return table
}

func assignCodePathToChars(root *Node, codePath string, table map[rune]string) {
	if root == nil {
		return
	}

	if root.value != 0 {
		table[root.value] = codePath
		return
	}

	assignCodePathToChars(root.left, codePath+"0", table)
	assignCodePathToChars(root.right, codePath+"1", table)
}

// ToReverseLookup assigns runes to Huffman Code paths to generate a reverse lookup map.
func (h *Node) ToReverseLookup() map[string]rune {
	lookup := h.ToLookup()
	reverse := make(map[string]rune, len(lookup))

	for char, codePath := range lookup {
		reverse[codePath] = char
	}

	return reverse
}
