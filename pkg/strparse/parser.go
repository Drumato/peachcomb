package strparse

type Parser[O ParseOutput] interface {
	Parse(input ParseInput) (ParseInput, O, ParseError)
}

type ParseInput string
type ParseOutput interface{}
