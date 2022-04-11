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

// Branches initializes a parser that receives multiple syntax-rules and determine one of them.
// In almost cases, user can enumerate all syntax rules before starting parsing.
// so branches receives map. We recommend you to initialize the map at once and refer multiple times.
// It may be efficient.
// if no applicable rule exists in the rules, Branches() parser returns an error.
// if all of them are failed to parse, Branches() parser also returns an error.
func Branches[E comparable, O parser.ParseOutput](rules map[E]parser.Parser[E, O]) parser.Parser[E, O] {
	return &branchesParser[E, O]{rules}
}

// branchesParser is the actual implementation of Branches() parser.
type branchesParser[E comparable, O parser.ParseOutput] struct {
	rules map[E]parser.Parser[E, O]
}

// Parse implements parser.Parser[E comparable, O parser.ParseOutput] interface.
func (p *branchesParser[E, O]) Parse(input parser.ParseInput[E]) (parser.ParseInput[E], O, parser.ParseError) {
	var o O
	if len(input) == 0 {
		return input, o, &parser.NoLeftInputToParseError{}
	}

	e := input[0]
	sub, ok := p.rules[e]
	if !ok {
		return input, o, &parser.NoLeftInputToParseError{}
	}

	return sub.Parse(input)
}

// ApplicableRuleIsNotFoundError notifies all of given parsers don't match the head of the input.
type ApplicableRuleIsNotFoundError[E comparable] struct {
	// actual is the given element
	actual E
}

// Error implements error interface.
func (e *ApplicableRuleIsNotFoundError[E]) Error() string {
	return fmt.Sprintf("all of given parser cannot start parsing on '%v'", e.actual)
}
