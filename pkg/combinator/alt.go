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

import "github.com/Drumato/goparsecomb/pkg/parser"

func Alt[E comparable, O parser.ParseOutput](parsers ...parser.Parser[E, O]) parser.Parser[E, O] {
	return &altParser[E, O]{parsers: parsers}
}

type altParser[E comparable, O parser.ParseOutput] struct {
	parsers []parser.Parser[E, O]
}

func (p *altParser[E, O]) Parse(input parser.ParseInput[E]) (parser.ParseInput[E], O, parser.ParseError) {
	var subI parser.ParseInput[E]
	var subO O
	var err parser.ParseError

	for _, subP := range p.parsers {
		subI, subO, err = subP.Parse(input)
		if err == nil {
			return subI, subO, nil
		}
	}

	return subI, subO, &AllParsersFailledError{}
}

// AllParsersFailledError notifies all of given parsers are failed to parse.
type AllParsersFailledError struct{}

// Error implements error interface.
func (e *AllParsersFailledError) Error() string {
	return "all of given parser failed to parse"
}
