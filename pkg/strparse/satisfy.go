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
	"fmt"

	"github.com/Drumato/goparsecomb/pkg/parser"
)

// Satisfy initializes a parser that checks the head of the input satisfies the predicate.
func Satisfy(pred Predicate) parser.Parser[string, rune] {
	return &satisfyParser{
		pred: pred,
	}
}

// satisfyParser is the actual implementation of Parser interface
type satisfyParser struct {
	pred Predicate
}

var _ parser.Parser[string, rune] = &satisfyParser{}

// Predicate is the condition that satisfyParser uses for consuming one rune.
type Predicate func(ch rune) bool

// Parse implements Parser[rune] interface
func (p *satisfyParser) Parse(input string) (string, rune, parser.ParseError) {
	if len(input) == 0 {
		return input, 0, &parser.NoLeftInputToParseError[string]{}
	}

	ch := []rune(input)[0]
	notSatisfied := !p.pred(ch)
	if notSatisfied {
		return input, 0, &NotSatisfiedError{}
	}

	// input[1:] doesn't split multi-byte string properly
	// so we should cast it into []rune first.
	rest := []rune(input)[1:]
	return string(rest), ch, nil
}

// NotsatisfiedError notifies that the given predicate is not satisfied.
type NotSatisfiedError struct {
	// actual is the given rune that satisfyParser consumed
	actual rune
}

var _ parser.ParseError = &NotSatisfiedError{}

// Error implements error interface
func (e *NotSatisfiedError) Error() string {
	return fmt.Sprintf("predicate was not satisfied on '%c'", e.actual)
}
