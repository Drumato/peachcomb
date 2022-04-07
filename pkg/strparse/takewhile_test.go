package strparse_test

import (
	"fmt"
	"testing"

	"github.com/Drumato/goparsecomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func Takewhile1(t *testing.T) {
	p := strparse.TakeWhile1(strparse.Satisfy(func(ch rune) bool {
		return ch == 'a'
	}))

	i, o, err := p.Parse(strparse.ParseInput("aaaabaa"))
	assert.NoError(t, err)
	assert.Equal(t, "aaaa", o)
	assert.Equal(t, strparse.ParseInput("baa"), i)
}
