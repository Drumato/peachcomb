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

package combinator_test

import (
	"fmt"

	"github.com/Drumato/peachcomb/pkg/byteparse"
	"github.com/Drumato/peachcomb/pkg/combinator"
	"github.com/Drumato/peachcomb/pkg/parser"
	"github.com/Drumato/peachcomb/pkg/strparse"
)

func ExamplePreceded() {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinator.Preceded(predecessor, successor)

	i, o, err := p([]rune("*a"))
	fmt.Println(string(i))
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	//
	// a
	// <nil>
}

func ExampleTerminated() {
	predecessor := strparse.Rune('a')
	successor := strparse.Rune('+')
	p := combinator.Terminated(predecessor, successor)

	i, o, err := p([]rune("a+"))
	fmt.Println(string(i))
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	//
	// a
	// <nil>
}

func ExampleSeparated1() {
	element := strparse.Digit1()
	separator := strparse.Rune('|')
	p := combinator.Separated1(element, separator)
	i, o, err := p([]rune("123|456|789Drumato"))
	fmt.Println(string(i))
	fmt.Printf("%d\n", len(o))
	fmt.Printf("%s %s %s\n", o[0], o[1], o[2])
	fmt.Println(err)
	// Output:
	// Drumato
	// 3
	// 123 456 789
	// <nil>
}

func ExampleSatisfy() {
	i, o, err := combinator.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})([]rune("abc"))
	fmt.Println(string(i))
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	//
	// bc
	// a
	// <nil>
}

func ExampleMap() {
	subsubP := strparse.Rune('a')
	subP := combinator.Many1(subsubP)
	p := combinator.Map(subP, func(s []rune) (int, error) { return len(s), nil })
	i, o, err := p([]rune("aaaabaaaa"))
	fmt.Println(string(i))
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	// baaaa
	// 4
	// <nil>
}

func ExampleAlt() {
	p1 := strparse.Rune('a')
	p2 := strparse.Rune('b')
	p := combinator.Many1(combinator.Alt(p1, p2))

	i, o, err := p([]rune("abababc"))
	fmt.Println(string(i))
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// c
	// ababab
	// <nil>
}

func ExampleDelimited() {
	begin := strparse.Rune('(')
	end := strparse.Rune(')')
	contents := strparse.Digit1()
	p := combinator.Delimited(begin, contents, end)

	i, o, err := p([]rune("(12321)"))
	fmt.Println(string(i))
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// 12321
	// <nil>
}

func ExampleMany0() {
	p := combinator.Many0(strparse.Rune('a'))

	i, o, err := p([]rune("baaaa"))
	fmt.Println(string(i))
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// baaaa
	//
	// <nil>
}

func ExampleMany1() {
	p := combinator.Many1(strparse.Rune('a'))

	i, o, err := p([]rune("aaaabaa"))
	fmt.Println(string(i))
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// baa
	// aaaa
	// <nil>
}

func ExampleManyMinMax() {
	p := combinator.ManyMinMax(strparse.Rune('a'), 3, 5)

	i, o, err := p([]rune("aaaabbb"))
	fmt.Println(string(i))
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// bbb
	// aaaa
	// <nil>
}

func ExampleBranches() {
	m := make(map[byte]parser.Parser[byte, string])
	m[0x00] = combinator.Map(byteparse.UInt8(), func(v uint8) (string, error) { return "0x00", nil })
	m[0x01] = combinator.Map(byteparse.UInt8(), func(v uint8) (string, error) { return "0x01", nil })

	p := combinator.Many1(combinator.Branches(m))
	i, o, err := p([]byte{0x00, 0x01, 0x00, 0x01, 0x02})
	fmt.Println(i)
	fmt.Println(len(o))
	fmt.Printf("%s %s %s %s\n", o[0], o[1], o[2], o[3])
	fmt.Println(err)
	// Output:
	// [2]
	// 4
	// 0x00 0x01 0x00 0x01
	// <nil>
}
