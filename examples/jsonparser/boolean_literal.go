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

package main

import (
	"fmt"

	"github.com/Drumato/peachcomb/pkg/byteparse"
	"github.com/Drumato/peachcomb/pkg/combinator"
	"github.com/Drumato/peachcomb/pkg/parser"
)

type jsonBooleanValue bool

// parseJSONBooleanValue parses the boolean literal.
// boolean := true | false
func parseJSONBooleanValue(input parser.ParseInput[byte]) (parser.ParseInput[byte], jsonValue, parser.ParseError) {
	const trueSig = "true"
	const falseSig = "false"
	trueP := byteparse.Tag([]byte(trueSig))
	falseP := byteparse.Tag([]byte(falseSig))
	p := combinator.Map(combinator.Alt(trueP, falseP), func(s []byte) (jsonValue, error) {
		switch string(s) {
		case trueSig:
			return jsonBooleanValue(true), nil
		case falseSig:
			return jsonBooleanValue(false), nil
		default:
			// maybe unreachable
			return nil, fmt.Errorf("unexpected bytes tag '%s'", s)
		}
	})

	return p(input)
}
