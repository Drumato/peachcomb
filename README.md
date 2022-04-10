# peachcomb

[![Go Reference](https://pkg.go.dev/badge/github.com/Drumato/peachcomb.svg)](https://pkg.go.dev/github.com/Drumato/peachcomb)  

Go Parser Combinator with Go Generics.  

## How to use

currently this library is so simple. you only should follow 2 steps as below.

- initializes `parser.Parser[I, O]`
- call `func (p Parser[I, O]) Parse(input I) (I, O, parser.ParseError)` method

## Example

```go
package main

import (
    "github.com/Drumato/peachcomb/pkg/strparse"
    "github.com/Drumato/peachcomb/pkg/combinator"
)

func main() {
	element := strparse.Digit1()
	separator := strparse.Rune('|')
	p := combinator.Separated1(element, separator)
	i, o, err := p.Parse([]rune("123|456|789Drumato"))
	fmt.Println(i)
	fmt.Printf("%d\n", o)
	fmt.Printf("%s %s %s\n", o[0], o[1], o[2])
	fmt.Println(err)
}
```

```shell
$ go run main.go
Drumato
123 456 789
3
<nil>
```
