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

package combinator

import (
	"github.com/Drumato/peachcomb/pkg/parser"
)

// Terminated initializes a parser that applies given parsers but discards successor's output.
func Terminated[
	E comparable,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
](predecessor parser.Parser[E, O1], successor parser.Parser[E, O2]) parser.Parser[E, O1] {
	return func(input parser.ParseInput[E]) (parser.ParseInput[E], O1, parser.ParseError) {

		rest, o1, err := predecessor(input)
		if err != nil {
			return rest, o1, err
		}

		rest, _, err = successor(rest)
		return rest, o1, err
	}
}
