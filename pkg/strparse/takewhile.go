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
	"strings"

	"github.com/Drumato/goparsecomb/pkg/parser"
)

// takeWhile1Parser is the actual implementation of Parser interface
type takeWhile1Parser struct {
	sub parser.Parser[string, rune]
}

// TakeWhile1 initializes a parser that applies the given sub-parser several times.
// if the sub parser fails to parse and the count of application times is 0
// TakeWhile1 parser return an error.
func TakeWhile1(sub parser.Parser[string, rune]) parser.Parser[string, string] {
	return &takeWhile1Parser{sub: sub}
}

// Parse implements Parser[string] interface
func (p *takeWhile1Parser) Parse(input string) (string, string, parser.ParseError) {
	if len(input) == 0 {
		return input, "", &parser.NoLeftInputToParseError[string]{}
	}

	count := 0
	var subI string
	var subO rune
	var subErr error
	var output strings.Builder
	for {
		subI, subO, subErr = p.sub.Parse(input[count:])
		if subErr != nil {
			break
		}
		count++

		output.WriteRune(subO)
	}

	if count == 0 {
		return subI, output.String(), subErr
	}

	return subI, output.String(), nil
}
