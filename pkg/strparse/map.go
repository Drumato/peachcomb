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
	"github.com/Drumato/goparsecomb/pkg/parser"
)

type mapParser[SO parser.ParseOutput, O parser.ParseOutput] struct {
	sub parser.Parser[string, SO]
	fn  func(SO) O
}

func Map[SO parser.ParseOutput, O parser.ParseOutput](sub parser.Parser[string, SO], fn func(SO) O) parser.Parser[string, O] {
	return &mapParser[SO, O]{sub: sub, fn: fn}
}

func (p *mapParser[SO, O]) Parse(input string) (string, O, parser.ParseError) {
	i, o, err := p.sub.Parse(input)
	if err != nil {
		return i, p.fn(o), err
	}

	return i, p.fn(o), err
}
