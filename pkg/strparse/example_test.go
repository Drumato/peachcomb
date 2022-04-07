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

package strparse_test

import (
	"fmt"

	"github.com/Drumato/goparsecomb/pkg/strparse"
)

func ExampleMap() {
	subsubP := strparse.Rune('a')
	subP := strparse.TakeWhile1(subsubP)
	p := strparse.Map(subP, func(s string) int { return len(s) })
	i, o, err := p.Parse("aaaabaaaa")
	fmt.Println(i)
	fmt.Printf("%d\n", o)
	fmt.Println(err)

	// Output:
	//
	// baaaa
	// 4
	// <nil>
}

func ExampleTakeWhile1() {
	p := strparse.TakeWhile1(strparse.Rune('a'))

	i, o, err := p.Parse("aaaabaa")
	fmt.Println(i)
	fmt.Println(o)
	fmt.Println(err)

	// Output:
	// baa
	// aaaa
	// <nil>
}

func ExampleSatisfy() {
	i, o, err := strparse.Satisfy(func(ch rune) bool {
		return ch == 'a'
	}).Parse("abc")
	fmt.Println(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)

	// Output:
	//
	// bc
	// a
	// <nil>
}

func ExampleRune() {
	i, o, err := strparse.Rune('a').Parse("abc")
	fmt.Println(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)

	// Output:
	//
	// bc
	// a
	// <nil>
}
