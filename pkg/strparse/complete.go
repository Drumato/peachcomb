// MIT License
//
// Copyright (c) 2022 Drumato
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package strparse

import (
	"github.com/Drumato/peachcomb/pkg/parser"
)

// CompleteInput holds the whole runes.
type CompleteInput struct {
	runes  []rune
	offset int
}

// NewCompleteInput initialiizes a CompleteInput.
func NewCompleteInput(s string) *CompleteInput {
	return &CompleteInput{runes: []rune(s)}
}

// Read implements parser.ParseInput interface.
func (c *CompleteInput) Read(buf []rune) (int, error) {
	if c.offset >= len(c.runes) {
		return 0, &parser.NoLeftInputToParseError{}
	}

	copy(buf, c.runes[c.offset:])
	c.offset += len(buf)

	return len(buf), nil
}

// Seek implements parser.ParseInput interface.
func (c *CompleteInput) Seek(n int, mode parser.SeekMode) (int, error) {
	switch mode {
	case parser.SeekModeStart:
		if n >= len(c.runes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset = n
		return c.offset, nil
	case parser.SeekModeCurrent:
		if c.offset+n >= len(c.runes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset += n
		return c.offset, nil
	default:
		panic("given seek mode is not supported")
	}
}
