package huffman

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestCountCharsFrequency(t *testing.T) {
	tests := []struct {
		input string
		want  map[rune]int
	}{
		{input: "", want: map[rune]int{}},
		{input: "test", want: map[rune]int{116: 2, 101: 1, 115: 1}},
		{input: "test test", want: map[rune]int{116: 4, 101: 2, 115: 2, 32: 1}},
		{input: "test\ntest", want: map[rune]int{116: 4, 101: 2, 115: 2, 10: 1}},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		got := getRunesFrequency(input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestWriteHeader(t *testing.T) {
	var buff bytes.Buffer
	want := struct {
		length string
		input  string
	}{
		length: "10",
		input:  "serialized",
	}

	err := writeHeader(want.input, &buff)
	if err != nil {
		t.Errorf("got unexpected error with writeHeader(%s, %v): %v", want.input, buff, err)
	}

	got := strings.Split(string(buff.Bytes()), EndOfHeaderLength)

	if got[0] != want.length {
		t.Errorf("got %s, want %s", got[0], want.length)
	}

	if got[1] != want.input {
		t.Errorf("got %s, want %s", got[1], want.input)
	}
}

func TestCompress(t *testing.T) {
	defer removeFiles(compressed, "")

	err := NewCompressor(original, compressed).Compress()
	if err != nil {
		t.Errorf("got unexpected error with NewCompressor(%s, %s).Compress(): %v", original, compressed, err)
	}

	numOfBytes, sameContent := haveSameContent(original, compressed)
	if sameContent {
		t.Errorf("compressed content is equal to original file")
	}

	if numOfBytes[0] >= numOfBytes[1] {
		t.Errorf("compressed content is bigger than original file")
	}
}

func BenchmarkCompress(b *testing.B) {
	b.StopTimer()
	defer removeFiles(compressed, "")
	c := NewCompressor(original, compressed)
	b.StartTimer()
	
	for i := 0; i < b.N; i++ {
		c.Compress()
	}
}
