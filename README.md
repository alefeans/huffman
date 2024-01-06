# huffman

`huffman` is a Go implementation of the [Huffman Coding](https://en.wikipedia.org/wiki/Huffman_coding) lossless data compression algorithm and the solution for the challenge [Write Your Own Compression Tool](https://codingchallenges.fyi/challenges/challenge-huffman/).

### Usage

```sh
go build

./huffman -h
Usage of ./huffman:
  -d    If set, input file will be decompressed to the output file
  -f string
        The name of the input file to be compressed
  -o string
        The name of the output compressed file
```

To compress a file:

```sh
./huffman -f test.txt -o compressed.txt

ls -lh test.txt compressed.txt
-rw-r--r--  1 alefeans  staff   1.8M Jan  3 19:40 compressed.txt
-rw-r--r--@ 1 alefeans  staff   3.1M Jan  2 17:44 test.txt
```

To decompress a file:

```sh
./huffman -f compressed.txt -o original.txt -d


ls -lh *.txt
-rw-r--r--  1 alefeans  staff   1.8M Jan  3 19:40 compressed.txt
-rw-r--r--  1 alefeans  staff   3.1M Jan  3 19:42 original.txt
-rw-r--r--@ 1 alefeans  staff   3.1M Jan  2 17:44 test.txt

diff test.txt original.txt
```

### Tests

```sh
go test -v ./...
?       github.com/alefeans/huffman     [no test files]
?       github.com/alefeans/huffman/internal/bit        [no test files]
=== RUN   TestCountCharsFrequency
--- PASS: TestCountCharsFrequency (0.00s)
=== RUN   TestWriteHeader
--- PASS: TestWriteHeader (0.00s)
=== RUN   TestCompress
--- PASS: TestCompress (0.03s)
=== RUN   TestReadHeader
--- PASS: TestReadHeader (0.00s)
=== RUN   TestDecompress
--- PASS: TestDecompress (0.04s)
=== RUN   TestPriorityQueue
--- PASS: TestPriorityQueue (0.00s)
=== RUN   TestNewHuffmanTree
--- PASS: TestNewHuffmanTree (0.00s)
=== RUN   TestToHeader
--- PASS: TestToHeader (0.00s)
=== RUN   TestFromHeader
--- PASS: TestFromHeader (0.00s)
=== RUN   TestToAndFromHeader
--- PASS: TestToAndFromHeader (0.00s)
=== RUN   TestToLookup
--- PASS: TestToLookup (0.00s)
=== RUN   TestToReverseLookup
--- PASS: TestToReverseLookup (0.00s)
PASS
ok      github.com/alefeans/huffman/internal/huffman    0.221s
```
### Benchmarks

```sh
go test ./... -bench=. -benchmem
?       github.com/alefeans/huffman     [no test files]
?       github.com/alefeans/huffman/internal/bit        [no test files]
goos: darwin
goarch: arm64
pkg: github.com/alefeans/huffman/internal/huffman
BenchmarkCompress-10                  57          18243634 ns/op           48923 B/op             461 allocs/op
BenchmarkDecompress-10               130           9151706 ns/op          405467 B/op           81198 allocs/op
PASS
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
