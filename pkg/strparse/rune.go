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

// Rune initializes a parser that consumes one rune.
// expected is the expected rune that you want to consume
func Rune(expected rune) parser.Parser[rune, rune] {
	return &runeParser{
		expected: expected,
	}
}

// runeParser is the actual impelementation of Parser interface
type runeParser struct {
	expected rune
}

// Parse implements Parser[string, rune] interface
func (p *runeParser) Parse(input parser.ParseInput[rune]) (parser.ParseInput[rune], rune, parser.ParseError) {
	if len(input) == 0 {
		return input, 0, &parser.NoLeftInputToParseError{}
	}

	ch := input[0]
	matched := ch == p.expected

	if !matched {
		return input, 0, &UnexpectedRuneError{actual: ch, expected: p.expected}
	}

	return input[1:], p.expected, nil
}

// UnexpectedRuneError notifies the head of the given input is unexpected.
type UnexpectedRuneError struct {
	actual   rune
	expected rune
}

// Error implements error interface.
func (e *UnexpectedRuneError) Error() string {
	return fmt.Sprintf("expected '%c' but got '%c'", e.expected, e.actual)
}
