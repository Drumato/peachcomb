package strparse_test

import (
	"testing"

	"github.com/Drumato/goparsecomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	subP := strparse.Rune('a')
	p := strparse.Map(subP, func(ch rune) bool { return ch == 'a' })

	i, o, err := p.Parse("abc")
	assert.NoError(t, err)
	assert.Equal(t, "bc", i)
	assert.Equal(t, true, o)
}
