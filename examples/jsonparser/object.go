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
	"github.com/Drumato/peachcomb/pkg/byteparse"
	"github.com/Drumato/peachcomb/pkg/combinator"
	"github.com/Drumato/peachcomb/pkg/parser"
)

type jsonObjectValue struct {
}

type jsonObjectField struct {
	Name  string
	Value jsonValue
}

// parseJSONObjectField parses a field of the JSON object
// object_field := string whitespace ":" value
func parseJSONObjectField(input parser.ParseInput[byte]) (parser.ParseInput[byte], jsonObjectField, error) {
	rawName := parseJSONStringValue
	ws := parseJSONWhitespace
	nameWS := combinator.Terminated(rawName, ws)
	colon := byteparse.Tag([]byte(":"))
	nameP := combinator.Terminated(nameWS, colon)
	p := combinator.Twin(nameP, parseJSONValue)

	i, o, err := p(input)
	if err != nil {
		return i, jsonObjectField{}, err
	}

	name, _ := o.One.(jsonStringValue)
	f := jsonObjectField{
		Name:  string(name),
		Value: o.Two,
	}
	return i, f, nil
}
