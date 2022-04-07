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

// Rune initializes a parser that consumes one rune.
// expected is the expected rune that you want to consume
func Rune(expected rune) Parser[rune] {
	return &runeParser{
		expected: expected,
	}
}

// runeParser is the actual impelementation of Parser interface
type runeParser struct {
	expected rune
}

var _ Parser[rune] = &runeParser{}

// Parse implements Parser[rune] interface
func (p *runeParser) Parse(input ParseInput) (ParseInput, rune, ParseError) {
	if len(input) == 0 {
		return input, 0, &NoLeftInputToParseError{}
	}

	ch := []rune(input)[0]
	matched := ch == p.expected

	if !matched {
		return input, 0, &UnexpectedRuneError{actual: ch, expected: p.expected}
	}

	// input[1:] doesn't split multi-byte string properly
	// so we should cast it into []rune first.
	rest := []rune(input)[1:]
	return ParseInput(rest), p.expected, nil
}
