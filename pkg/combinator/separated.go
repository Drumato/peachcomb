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
	"github.com/Drumato/peachcomb/pkg/parser"
)

func Separated1[E comparable, EO parser.ParseOutput, SO parser.ParseOutput](element parser.Parser[E, EO], separator parser.Parser[E, SO]) parser.Parser[E, []EO] {
	return func(input parser.ParseInput[E]) (parser.ParseInput[E], []EO, parser.ParseError) {
		output := make([]EO, 0)
		rest, e1, err := element(input)
		if err != nil {
			return rest, output, err
		}

		output = append(output, e1)

		for {
			var eo EO

			// we mustn't generate an error if separator parser fails
			// because such as case is the end of the separated-list.
			rest, _, err = separator(rest)
			if err != nil {
				break
			}

			rest, eo, err = element(rest)
			if err != nil {
				// must generate an error.
				return rest, output, err
			}

			output = append(output, eo)
		}

		// we mustn't return err if separator parser fails
		// because such as case is the end of the separated-list.
		return rest, output, nil
	}
}
