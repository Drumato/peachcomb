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
	"strings"

	"github.com/Drumato/goparsecomb/pkg/parser"
)

// Tag initializes a parser that checks the input starts with the tag prefix.
func Tag(tag string) parser.Parser[string, string] {
	return &tagParser{
		tag: tag,
	}
}

// tagParser is the actual implementation of Parser interface
type tagParser struct {
	tag string
}

// Parse implements Parser[string, string] interface
func (p *tagParser) Parse(input string) (string, string, parser.ParseError) {
	if len(input) < len(p.tag) {
		return input, "", &parser.NoLeftInputToParseError[string]{}
	}

	unmatched := !strings.HasPrefix(input, p.tag)
	if unmatched {
		return input, "", &UnexpectedPrefixError{expected: p.tag}
	}

	return strings.TrimPrefix(input, p.tag), p.tag, nil
}

// UnexpectedPrefixError notifies the prefix of the given input is unexpected.
type UnexpectedPrefixError struct {
	expected string
}

// Error implements error interface.
func (e *UnexpectedPrefixError) Error() string {
	return fmt.Sprintf("expected \"%s\" prefix", e.expected)
}
