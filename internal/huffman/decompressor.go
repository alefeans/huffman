package huffman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/alefeans/huffman/internal/bit"
)

type Decompressor struct {
	inputFile  string
	outputFile string
}

func NewDecompressor(inputFile, outputFile string) *Decompressor {
	return &Decompressor{
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}

func (d *Decompressor) Decompress() error {
	input, err := os.Open(d.inputFile)
	if err != nil {
		return fmt.Errorf("can't open input file %s: %v\n", d.inputFile, err)
	}
	defer input.Close()

	output, err := os.Create(d.outputFile)
	if err != nil {
		return fmt.Errorf("can't create output file %s: %v\n", d.outputFile, err)
	}
	defer output.Close()

	header, err := readHeader(input)
	if err != nil {
		return fmt.Errorf("can't deserialize header: %v\n", err)
	}

	tree, err := FromHeader(header)
	if err != nil {
		return fmt.Errorf("can't deserialize header: %v\n", err)
	}

	return writeDecompressedContent(tree.ToReverseLookup(), output, bit.NewReader(input))
}

// readHeader first reads the length of the header from input to know where it ends,
// then it reads the whole header in one go using the length size.
func readHeader(input io.Reader) (string, error) {
	var headerLength string
	char := make([]byte, 1)

	for {
		_, err := input.Read(char)
		if err != nil {
			return "", err
		}

		token := string(char)
		if token == EndOfHeaderLength {
			break
		}

		headerLength += token
	}

	buffLength, err := strconv.Atoi(headerLength)
	if err != nil {
		return "", err
	}

	header := make([]byte, buffLength)
	if _, err := io.ReadFull(input, header); err != nil {
		return "", err
	}

	return string(header), nil
}

// writeDecompressedContent writes to output one rune at a time based on the lookup map.
func writeDecompressedContent(lookup map[string]rune, output io.Writer, br *bit.Reader) error {
	code := ""
	writer := bufio.NewWriter(output)

	for {
		b, err := br.ReadBit()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if b == bit.BitZero {
			code += "0"
		} else {
			code += "1"
		}

		if val, ok := lookup[code]; ok {
			_, err := writer.WriteRune(val)
			if err != nil {
				return err
			}
			code = ""
		}
	}

	writer.Flush()
	return nil
}
