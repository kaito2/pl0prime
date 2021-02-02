package reader

import (
	"bufio"
	"io"
	"log"

	"golang.org/x/xerrors"
)

type CharReader interface {
	NextChar() (rune, bool, error)
}

type charReader struct {
	input     *bufio.Scanner
	output    *bufio.Writer
	line      []rune
	lineIndex int
}

func (c *charReader) NextChar() (rune, bool, error) {
	if c.lineIndex == len(c.line) {
		c.lineIndex = -1
		return ' ', true, nil
	}

	if c.lineIndex == -1 {
		if c.input.Scan() {
			// TODO: Convert []byte to []rune directory.
			c.line = []rune(string(c.input.Bytes()))

			// TODO: remove debug print
			log.Printf("read line: %#v (%d)", c.line, len(c.line))

			c.lineIndex = 0
		} else {
			if err := c.input.Err(); err != nil {
				// TODO: return nil value instead of hard coded '0'
				return '0', false, xerrors.Errorf("failed to scan file: %w", err)
			}
			return '0', false, nil
		}
	}
	ch := c.line[c.lineIndex]
	c.lineIndex++
	return ch, true, nil
}

func NewCharReader(input io.Reader, output io.Writer) CharReader {
	scanner := bufio.NewScanner(input)
	bufioOutput := bufio.NewWriter(output)

	return &charReader{
		input:     scanner,
		output:    bufioOutput,
		line:      nil,
		lineIndex: -1,
	}
}
