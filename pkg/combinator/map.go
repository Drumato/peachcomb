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
	"github.com/Drumato/goparsecomb/pkg/parser"
)

type mapParser[I parser.ParseInput, SO parser.ParseOutput, O parser.ParseOutput] struct {
	sub parser.Parser[I, SO]
	fn  func(SO) (O, error)
}

func Map[I parser.ParseInput, SO parser.ParseOutput, O parser.ParseOutput](sub parser.Parser[I, SO], fn func(SO) (O, error)) parser.Parser[I, O] {
	return &mapParser[I, SO, O]{sub: sub, fn: fn}
}

func (p *mapParser[I, SO, O]) Parse(input I) (I, O, parser.ParseError) {
	var o O

	i, so, err := p.sub.Parse(input)
	if err != nil {
		return i, o, err
	}

	o, err = p.fn(so)
	return i, o, err
}
