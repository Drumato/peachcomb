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

// Uint32 initializes a parser that parse 32-bit unsigned integer.
// user can determine the behavior of this parser by giving byteorder what you want to use.
func UInt32(byteorder binary.ByteOrder) parser.Parser[byte, uint32] {
	return func(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint32, parser.ParseError) {
		buf := make([]byte, 4)

		n, err := input.Read(buf)
		if err != nil || n < 4 {
			return input, 0, &parser.NoLeftInputToParseError{}
		}

		v := byteorder.Uint32(buf)
		return input, v, nil
	}
}
