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

func Delimited[
	I parser.ParseInput,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
	O3 parser.ParseOutput,
](
	begin parser.Parser[I, O1],
	contents parser.Parser[I, O2],
	end parser.Parser[I, O3],
) parser.Parser[I, O2] {
	return &delimitedParser[I, O1, O2, O3]{begin: begin, contents: contents, end: end}
}

type delimitedParser[
	I parser.ParseInput,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
	O3 parser.ParseOutput,
] struct {
	begin    parser.Parser[I, O1]
	contents parser.Parser[I, O2]
	end      parser.Parser[I, O3]
}

func (p *delimitedParser[I, O1, O2, O3]) Parse(input I) (I, O2, parser.ParseError) {
	var o2 O2
	rest, _, err := p.begin.Parse(input)
	if err != nil {
		return rest, o2, err
	}

	rest, o2, err = p.contents.Parse(rest)
	if err != nil {
		return rest, o2, err
	}

	rest, _, err = p.end.Parse(rest)
	if err != nil {
		return rest, o2, err
	}

	return rest, o2, nil
}
