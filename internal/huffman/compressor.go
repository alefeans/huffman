package huffman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/alefeans/huffman/internal/bit"
)

const EndOfHeaderLength = ":"

type Compressor struct {
	inputFile  string
	outputFile string
}

func NewCompressor(inputFile, outputFile string) *Compressor {
	return &Compressor{
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}

func (c *Compressor) Compress() error {
	input, err := os.Open(c.inputFile)
	if err != nil {
		return fmt.Errorf("can't open input file %s: %v\n", c.inputFile, err)
	}
	defer input.Close()

	output, err := os.Create(c.outputFile)
	if err != nil {
		return fmt.Errorf("can't create output file %s: %v\n", c.outputFile, err)
	}
	defer output.Close()

	runesFrequency := getRunesFrequency(input)
	pq := NewPriorityQueue(runesFrequency)
	tree := NewHuffmanTree(pq)

	err = writeHeader(tree.ToHeader(), output)
	if err != nil {
		return fmt.Errorf("error writing header: %v", err)
	}

	input.Seek(0, 0) // Reset reader to the beginning of file
	err = writeCompressedContent(tree.ToLookup(), input, bit.NewWriter(output))
	if err != nil {
		return fmt.Errorf("error writing compressed content: %v", err)
	}

	return nil
}

// getRunesFrequency counts the number of times a rune was found in input.
func getRunesFrequency(input io.Reader) map[rune]int {
	frequency := make(map[rune]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		r, _ := utf8.DecodeRuneInString(scanner.Text())
		frequency[r]++
	}

	return frequency
}

// writeHeader writes the length of the header and the serialized tree
// separated by a control character to output. The length is useful to
// know where the header ends in a future read.
func writeHeader(serialized string, output io.Writer) error {
	header := fmt.Sprintf("%d%s%s", len(serialized), EndOfHeaderLength, serialized)
	_, err := output.Write([]byte(header))
	if err != nil {
		return err
	}

	return nil
}

// writeCompressedContent writes one bit at a time to bw.writer based on lookup code path.
func writeCompressedContent(lookup map[rune]string, input io.Reader, bw *bit.Writer) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		r, _ := utf8.DecodeRuneInString(scanner.Text())

		for _, b := range lookup[r] {
			var err error
			if b == '0' {
				err = bw.Write(bit.BitZero)
			} else {
				err = bw.Write(bit.BitOne)
			}
			if err != nil {
				return err
			}
		}
	}

	bw.Flush()
	return nil
}
