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

import "github.com/Drumato/peachcomb/pkg/parser"

// Delimited initializes a parser that parses a delimited sequence. (e.g. '[foobar]')
func Delimited[
	E comparable,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
	O3 parser.ParseOutput,
](
	begin parser.Parser[E, O1],
	contents parser.Parser[E, O2],
	end parser.Parser[E, O3],
) parser.Parser[E, O2] {
	return func(input parser.ParseInput[E]) (parser.ParseInput[E], O2, parser.ParseError) {
		var o2 O2
		rest, _, err := begin(input)
		if err != nil {
			return rest, o2, err
		}

		rest, o2, err = contents(rest)
		if err != nil {
			return rest, o2, err
		}

		rest, _, err = end(rest)
		if err != nil {
			return rest, o2, err
		}

		return rest, o2, nil
	}
}
