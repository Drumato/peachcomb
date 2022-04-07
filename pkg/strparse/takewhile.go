package strparse

import "strings"

type takeWhile1Parser struct {
	sub Parser[rune]
}

var _ Parser[string] = &takeWhile1Parser{}

func TakeWhile1(sub Parser[rune]) Parser[string] {
	return &takeWhile1Parser{sub: sub}
}

func (p *takeWhile1Parser) Parse(input ParseInput) (ParseInput, string, ParseError) {
	if len(input) == 0 {
		return input, "", &NoLeftInputToParseError{}
	}

	count := 0
	var subI ParseInput
	var subO rune
	var subErr error
	var output strings.Builder
	for {
		subI, subO, subErr = p.sub.Parse(input[count:])
		if subErr != nil {
			break
		}
		count++

		output.WriteRune(subO)
	}

	if count == 0 {
		return subI, output.String(), subErr
	}

	return subI, output.String(), nil
}
