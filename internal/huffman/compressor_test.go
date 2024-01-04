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
	input := "serialized"

	err := writeHeader(input, &buff)
	if err != nil {
		t.Errorf("got unexpected error with writeHeader(%s, %v)", input, buff)
	}
}
