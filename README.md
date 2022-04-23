# peachcomb

[![Go Reference](https://pkg.go.dev/badge/github.com/Drumato/peachcomb.svg)](https://pkg.go.dev/github.com/Drumato/peachcomb)  

Go Parser Combinator with Go Generics.  

## How to use

currently this library is so simple. you only should follow 2 steps as below.

- initializes `parser.Parser[I, O]`
- call the parser

## Examples

### (WIP) JSON Parser

```shell
$ go run ./examples/jsonparser/ examples/jsonparser/example.json
{[{[a b] 2} {[c d] 2} {[e f] 2}] 3}
```

### Simplest Case

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

	i := strparse.NewCompleteInput("123|456|789Drumato")
	_, o, err := p(i)
	fmt.Printf("%d\n", len(o))
	fmt.Printf("%s %s %s\n", o[0], o[1], o[2])
	fmt.Println(err)
}
```

```shell
$ go run main.go
3
123 456 789
<nil>
```
