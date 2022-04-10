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
	"os"

	"github.com/Drumato/goparsecomb/pkg/combinator"
	"github.com/Drumato/goparsecomb/pkg/parser"
)

type myTokenKind string

const (
	myTokenKindInteger myTokenKind = "<Integer>"
	myTokenKindLParen  myTokenKind = "("
	myTokenKindRParen  myTokenKind = ")"
)

type myToken struct {
	kind  myTokenKind
	p     position
	value int
}

// String() implemlents fmt.Stringer interface
func (t myToken) String() string {
	return fmt.Sprintf("(%d:%d) kind: %+v, value: %d", t.p.line, t.p.column, t.kind, t.value)
}

type position struct {
	line   uint
	column uint
}

func main() {
	p := setupParser()
	succeedCase(p)
	failCase(p)
}

func succeedCase(p parser.Parser[myToken, int]) {
	tokens := []myToken{
		// "("
		{kind: myTokenKindLParen, p: position{line: 1, column: 1}},
		// 12345
		{kind: myTokenKindInteger, p: position{line: 1, column: 2}, value: 12345},
		// ")"
		{kind: myTokenKindRParen, p: position{line: 1, column: 7}},
	}

	_, v, err := p.Parse(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(v)
}

func failCase(p parser.Parser[myToken, int]) {
	tokens := []myToken{
		// "("
		{kind: myTokenKindLParen, p: position{line: 1, column: 1}},
		// 12345
		{kind: myTokenKindInteger, p: position{line: 1, column: 2}, value: 12345},
		// 678910
		{kind: myTokenKindInteger, p: position{line: 1, column: 7}, value: 678910},
	}

	_, _, err := p.Parse(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	}
}

// parser := "(" integer-token ")"
func setupParser() parser.Parser[myToken, int] {
	// begin := "("
	begin := combinator.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindLParen
	})

	// contents := integer-token
	contents := combinator.Map(combinator.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindInteger
	}), func(t myToken) (int, error) {
		return t.value, nil
	})

	// end := ")"
	end := combinator.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindRParen
	})
	return combinator.Delimited(begin, contents, end)
}
