package strparse_test

import (
	"fmt"

	"github.com/Drumato/goparsecomb/pkg/strparse"
)

func ExampleTakeWhile1() {
	var subParser strparse.Parser[rune] = strparse.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	var p strparse.Parser[string] = strparse.TakeWhile1(subParser)

	i, o, _ := p.Parse(strparse.ParseInput("aaaabaa"))
	fmt.Println(i)
	fmt.Println(o)

	// Output:
	// baa
	// aaaa
}
