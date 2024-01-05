package huffman

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadHeader(t *testing.T) {
	want := "abcdeabcde"
	header := fmt.Sprintf("%d%s%s", len(want), EndOfHeaderLength, want)
	input := bytes.NewReader([]byte(header))

	got, err := readHeader(input)
	if err != nil {
		t.Errorf("got unexpected error with readHeader(%v): %v", input, err)
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
