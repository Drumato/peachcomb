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

package byteparse

import (
	"encoding/binary"

	"github.com/Drumato/peachcomb/pkg/parser"
)

// Uint64 initializes a parser that parse 64-bit unsigned integer.
// user can determine the behavior of this parser by giving byteorder what you want to use.
func UInt64(byteorder binary.ByteOrder) parser.Parser[byte, uint64] {
	return &uint64Parser{byteorder: byteorder}
}

// uint64Parser is the actual implementation of Uint64().
type uint64Parser struct {
	byteorder binary.ByteOrder
}

var _ parser.Parser[byte, uint64] = &uint64Parser{}

// Parse implements parser.Parser[byte, uint64] interface.
func (p *uint64Parser) Parse(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint64, parser.ParseError) {
	if len(input) < 8 {
		return nil, 0, &parser.NoLeftInputToParseError{}
	}

	v := p.byteorder.Uint64(input)
	return input[8:], v, nil
}
