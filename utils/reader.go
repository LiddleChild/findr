package utils

import (
	"bufio"
	"errors"
	"io"
)

type Reader struct {
	br    *bufio.Reader
	error error
}

func NewReader(rd io.Reader) *Reader {
	return &Reader{
		br:    bufio.NewReader(rd),
		error: nil,
	}
}

func (reader *Reader) NextRune() rune {
	r, _, err := reader.br.ReadRune()
	reader.error = err

	return r
}

func (reader *Reader) IsEOF() bool {
	return errors.Is(reader.error, io.EOF)
}

func (reader *Reader) Error() error {
	return reader.error
}
