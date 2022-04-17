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
	"fmt"

	"github.com/Drumato/peachcomb/pkg/parser"
)

// Many0 initializes a parser that applies the given sub-parser several times.
func Many0[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO]) parser.Parser[E, []SO] {
	return many(sub, 0)
}

// Many1 initializes a parser that applies the given sub-parser several times.
// if the sub parser fails to parse and the count of application times is 0
// Many11 parser return an error.
func Many1[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO]) parser.Parser[E, []SO] {
	return many(sub, 1)
}

// many is the actual implementation of Many0/1.
func many[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO], min uint) parser.Parser[E, []SO] {
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

// NotSatisfiedCountError notifies the count of sub-parser success are not satisfied.
type NotSatisfiedCountError struct {
	expected int
}

// Error implements error interface.
func (e *NotSatisfiedCountError) Error() string {
	return fmt.Sprintf("not satisfied '%d' sub-parser succeeds", e.expected)
}
