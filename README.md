# goparsecomb

[![Go Reference](https://pkg.go.dev/badge/github.com/Drumato/goparsecomb.svg)](https://pkg.go.dev/github.com/Drumato/goparsecomb)  

Go Parser Combinator with Go Generics.  

## How to use

currently this library is so simple. you only should follow 2 steps as below.

- initializes `parser.Parser[I, O]`
- call `func (p Parser[I, O]) Parse(input I) (I, O, parser.ParseError)` method

## Example

```go
package main

import (
    "github.com/Drumato/goparsecomb/pkg/strparse"
)

func main() {
	subsubP := strparse.Rune('a')
	subP := strparse.TakeWhile1(subsubP)
	p := strparse.Map(subP, func(s string) int { return len(s) })
	i, o, err := p.Parse("aaaabaaaa")
	fmt.Println(i)
	fmt.Printf("%d\n", o)
	fmt.Println(err)
}
```

```shell
$ go run main.go
baaaa
4
<nil>
```
