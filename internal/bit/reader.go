package bit

import "io"

const LeftmostBitSetInByte byte = 0b10000000

type Reader struct {
	reader    io.Reader
	buffer    [1]byte
	alignment uint8
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: r,
	}
}

// ReadBit first call fills the buffer with a byte read from reader and sets the
// alignment to MaxAligmentInByte (i.e the buffer is full). Then, it checks
// if the leftmost bit is set in the buffer, returning BitOne or BitZero.
// It shift lefts the buffer in one bit, so the leftmost bit is discarded and 
// the next bit becomes the leftmost (ex: "1010 1000" becomes "0101 0000").
func (r *Reader) ReadBit() (Bit, error) {
	if r.alignment == 0 {
		if n, err := r.reader.Read(r.buffer[:]); n != 1 || (err != nil && err != io.EOF) {
			return BitZero, err
		}

		r.alignment = MaxAligmentInByte
	}

	r.alignment--

	bit := BitZero
	if r.isLeftmostBitSetInBuffer() {
		bit = BitOne
	}

	r.buffer[0] <<= 1
	return bit, nil
}

// isLeftmostBitSetInBuffer checks if the buffer contains the leftmost bit
// set. Ex: if buffer contains bits "1011 0100", then it returns true. If 
// the buffer contains the bits "0100 0010", it returns false.
func (r *Reader) isLeftmostBitSetInBuffer() bool {
	return (r.buffer[0] & LeftmostBitSetInByte) != 0
}
