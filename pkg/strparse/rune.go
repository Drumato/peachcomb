package strparse

func Rune(expected rune) Parser[rune] {
	return &runeParser{
		expected: expected,
	}
}

type runeParser struct {
	expected rune
}

var _ Parser[rune] = &runeParser{}

func (p *runeParser) Parse(input ParseInput) (ParseInput, rune, ParseError) {
	if len(input) == 0 {
		return input, 0, &NoLeftInputToParseError{}
	}

	ch := []rune(input)[0]
	matched := ch == p.expected

	if !matched {
		return input, 0, &UnexpectedRuneError{actual: ch, expected: p.expected}
	}

	return ParseInput([]rune(input)[1:]), p.expected, nil
}
