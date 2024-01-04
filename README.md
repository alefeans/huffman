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

TBD

### Benchmarks

TBD

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
