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

package parser

// Parser is an abstract parser that parses string.
// All parsers in package strparse implements this interface.
type Parser[I ParseInput, O ParseOutput] interface {
	// Parse parses the input and convert the consumed substr to O actually
	Parse(input I) (I, O, ParseError)
}

// ParseInput is the input of Parser interface.
type ParseInput interface {
	string | []byte
}

// ParseOutput is the actual type of the parser's output
type ParseOutput interface{}

// ParseError represents the error of parsers in package strparse
type ParseError interface {
	error
}

// ErrorIs checks the given error implements ParseError interface.
func ErrorIs[T ParseError](err error, ty T) bool {
	_, ok := err.(T)
	return ok
}

// NoLeftInputToParseError notifies the given input to parser is empty
type NoLeftInputToParseError struct{}

// Error implements error interface
func (e *NoLeftInputToParseError) Error() string {
	return "no left input to parse"
}
