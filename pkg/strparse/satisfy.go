package strparse

import "fmt"

type satisfyParser struct {
	pred Predicate
}

var _ Parser[rune] = &satisfyParser{}

func Satisfy(pred Predicate) Parser[rune] {
	return &satisfyParser{
		pred: pred,
	}
}

type Predicate func(ch rune) bool

func (p *satisfyParser) Parse(input ParseInput) (ParseInput, rune, ParseError) {
	if len(input) == 0 {
		return input, 0, &NoLeftInputToParseError{}
	}

	ch := []rune(input)[0]
	notSatisfied := !p.pred(ch)
	if notSatisfied {
		return input, 0, &NotSatisfiedError{}
	}

	return ParseInput([]rune(input)[1:]), ch, nil
}

type NotSatisfiedError struct {
	actual rune
}

var _ ParseError = &NotSatisfiedError{}

func (e *NotSatisfiedError) Error() string {
	return fmt.Sprintf("predicate was not satisfied on '%c'", e.actual)
}
