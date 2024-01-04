package bit

import (
	"io"
)

type Bit byte

const (
	BitZero              Bit  = 0
	BitOne               Bit  = 1
	MaxAligmentInByte    byte = 8
)

type Writer struct {
	writer   io.Writer
	buffer   [1]byte
	aligment uint8
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		writer:   w,
		buffer:   [1]byte{0},
		aligment: MaxAligmentInByte,
	}
}

// Write sets one bit at a time to the buffer based on alignment. Ex:
// On the first call, the buffer contains the bits "0000 0000" and aligment is 8.
// If Write is called with bit "1", the bit is shifted to the leftmost position
// so the buffer becomes "1000 0000" and alignment is decremented. Next call
// with bit "1", will shift it to the next position based on the aligment,
// so the buffer becomes "1100 0000". When alignment is 0 (i.e the buffer is full),
// the whole byte is written to writer.
func (w *Writer) Write(bit Bit) error {
	if bit != 0 {
		w.buffer[0] |= 1 << (w.aligment - 1)
	}

	w.aligment--

	if w.aligment == 0 {
		if _, err := w.writer.Write(w.buffer[:]); err != nil {
			return err
		}

		w.buffer[0] = 0
		w.aligment = MaxAligmentInByte
	}

	return nil
}

func (w *Writer) Flush() error {
	for w.aligment != MaxAligmentInByte {
		err := w.Write(BitOne)
		if err != nil {
			return err
		}
	}

	return nil
}
