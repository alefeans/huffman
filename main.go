package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alefeans/huffman/internal/huffman"
)

type Args struct {
	inputFile  string
	outputFile string
	decompress bool
}

func parseCliArgs() *Args {
	var args Args
	flag.StringVar(&args.inputFile, "f", "", "The name of the input file to be compressed")
	flag.StringVar(&args.outputFile, "o", "", "The name of the output compressed file")
	flag.BoolVar(&args.decompress, "d", false, "If set, input file will be decompressed to the output file")
	flag.Parse()
	return &args
}

func main() {
	args := parseCliArgs()

	if args.inputFile == "" || args.outputFile == "" {
		fmt.Println("argument is missing")
		flag.Usage()
		os.Exit(1)
	}

	if args.decompress {
		err := huffman.NewDecompressor(args.inputFile, args.outputFile).Decompress()
		if err != nil {
			fmt.Printf("can't decompress file: %v\n", err)
			os.Exit(1)
		}
	} else {
		err := huffman.NewCompressor(args.inputFile, args.outputFile).Compress()
		if err != nil {
			fmt.Printf("can't compress file: %v\n", err)
			os.Exit(1)
		}
	}
}
