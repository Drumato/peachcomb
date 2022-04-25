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

// Satisfy initializes a parser that checks the head of the input satisfies the predicate.
func Satisfy[E comparable](pred Predicate[E]) parser.Parser[E, E] {
	return func(input parser.ParseInput[E]) (parser.ParseInput[E], E, parser.ParseError) {
		var e E

		storedOffset, err := input.Seek(0, parser.SeekModeCurrent)
		if err != nil {
			return input, e, err
		}

		buf := make([]E, 1)
		n, err := input.Read(buf)
		if err != nil || n < 1 {
			return input, e, &parser.NoLeftInputToParseError{}
		}

		notSatisfied := !pred(buf[0])
		if notSatisfied {
			// recover the consumed head of the input stream.
			input.Seek(storedOffset, parser.SeekModeStart)
			return input, e, &NotSatisfiedError[E]{actual: e}
		}

		return input, buf[0], nil
	}
}

// Predicate is the condition that satisfyParser uses for consuming one element.
type Predicate[E comparable] func(element E) bool

// NotsatisfiedError notifies that the given predicate is not satisfied.
type NotSatisfiedError[E comparable] struct {
	// actual is the given element that satisfyParser consumed
	actual E
}

// Error implements error interface
func (e *NotSatisfiedError[E]) Error() string {
	return fmt.Sprintf("predicate was not satisfied on '%+v'", e.actual)
}
