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

// ManyMinMax initializes a parser that applies the given sub-parser several times.
// It fails if the sub parser does not succeed at least min times.
// It also fails if the sub parser does succeed over max times.
func ManyMinMax[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO], min uint, max uint) parser.Parser[E, []SO] {
	return manyMinMax(sub, min, max)
}

// manyMinMax is the actual implementation of ManyMinMax.
func manyMinMax[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO], min uint, max uint) parser.Parser[E, []SO] {
	return func(input parser.ParseInput[E]) (parser.ParseInput[E], []SO, parser.ParseError) {
		if len(input) == 0 {
			return input, nil, &parser.NoLeftInputToParseError{}
		}

		count := 0
		output := make([]SO, 0)
		rest := input
		for {
			var o SO
			var err error

			if count >= int(max) {
				return rest, output, &NotSatisfiedCountError{}
			}

			rest, o, err = sub(rest)
			if err != nil {
				break
			}
			count++

			output = append(output, o)
		}

		if count < int(min) {
			return rest, output, &NotSatisfiedCountError{}
		}

		return rest, output, nil

	}
}
