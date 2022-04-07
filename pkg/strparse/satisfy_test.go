package strparse_test

import (
	"testing"

	"github.com/Drumato/goparsecomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestSatisfy(t *testing.T) {
	p := strparse.Satisfy(func(ch rune) bool { return ch == 'a' })
	i, o, err := p.Parse(strparse.ParseInput("abc"))

	assert.NoError(t, err)
	assert.Equal(t, strparse.ParseInput("bc"), i)
	assert.Equal(t, 'a', o)
}
