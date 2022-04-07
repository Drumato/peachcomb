package strparse

import "fmt"

type ParseError interface {
	error
}

type NoLeftInputToParseError struct{}

var _ ParseError = &NoLeftInputToParseError{}

func (e *NoLeftInputToParseError) Error() string {
	return "no left input to parse"
}

func ErrorIs[T ParseError](err error) bool {
	_, ok := err.(T)
	return ok
}

type UnexpectedRuneError struct {
	actual   rune
	expected rune
}

var _ ParseError = &UnexpectedRuneError{}

func (e *UnexpectedRuneError) Error() string {
	return fmt.Sprintf("expected %c but got %c", e.expected, e.actual)
}
