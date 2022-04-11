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

// Alt initializes a parser that applies all given parsers.
// if all of them are failed to parse, Alt() parser also returns an error.
// otherwise Alt() succeeds to parse.
func Alt[E comparable, O parser.ParseOutput](parsers ...parser.Parser[E, O]) parser.Parser[E, O] {
	return &altParser[E, O]{parsers: parsers}
}

// altParser is the actual implementation of Alt() parser.
type altParser[E comparable, O parser.ParseOutput] struct {
	parsers []parser.Parser[E, O]
}

// Parse implements parser.Parser[E comparable, O parser.ParseOutput] interface.
func (p *altParser[E, O]) Parse(input parser.ParseInput[E]) (parser.ParseInput[E], O, parser.ParseError) {
	// subI holds the rest input in outer scope of for-statement.
	var subI parser.ParseInput[E]
	var subO O

	for _, subP := range p.parsers {
		var err parser.ParseError

		subI, subO, err = subP.Parse(input)
		if err == nil {
			return subI, subO, nil
		}
	}

	return subI, subO, &AllParsersFailedError{}
}

// AllParsersFailedError notifies all of given parsers are failed to parse.
type AllParsersFailedError struct{}

// Error implements error interface.
func (e *AllParsersFailedError) Error() string {
	return "all of given parser failed to parse"
}
