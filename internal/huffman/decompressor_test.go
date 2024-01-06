package huffman

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	original     = "testdata/test.txt"
	compressed   = "testdata/compressed.txt"
	decompressed = "testdata/decompressed.txt"
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

func haveSameContent(file1, file2 string) ([2]int, bool) {
	var numOfBytes [2]int

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	f1Scan := bufio.NewScanner(f1)
	f2Scan := bufio.NewScanner(f2)

	for f1Scan.Scan() {
		f2Scan.Scan()

		f1Bytes := f1Scan.Bytes()
		f2Bytes := f2Scan.Bytes()

		numOfBytes[0] += len(f1Bytes)
		numOfBytes[1] += len(f2Bytes)

		if !bytes.Equal(f1Bytes, f2Bytes) {
			return numOfBytes, false
		}
	}

	return numOfBytes, true
}

func removeFiles(file1, file2 string) error {
	err := os.Remove(file1)
	if err != nil {
		return err
	}

	err = os.Remove(file2)
	if err != nil {
		return err
	}

	return nil
}

func TestDecompress(t *testing.T) {
	defer removeFiles(compressed, decompressed)

	err := NewCompressor(original, compressed).Compress()
	if err != nil {
		t.Errorf("got unexpected error with NewCompressor(%s, %s).Compress(): %v", original, compressed, err)
	}

	err = NewDecompressor(compressed, decompressed).Decompress()
	if err != nil {
		t.Errorf("got unexpected error with NewDecompressor(%s, %s).Decompress(): %v", compressed, decompressed, err)
	}

	numOfBytes, sameContent := haveSameContent(original, decompressed)
	if !sameContent && numOfBytes[0] != numOfBytes[1] {
		t.Errorf("decompressed content is different from original file")
	}
}

func BenchmarkDecompress(b *testing.B) {
	b.StopTimer()
	defer removeFiles(compressed, decompressed)
	NewCompressor(original, compressed).Compress()
	d := NewDecompressor(compressed, decompressed)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		d.Decompress()
	}
}
