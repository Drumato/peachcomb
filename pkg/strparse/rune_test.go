package strparse_test

import (
	"testing"

	"github.com/Drumato/goparsecomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestRuneOnASCII(t *testing.T) {
	p := strparse.Rune('a')
	i, o, err := p.Parse("abc")
	assert.NoError(t, err)
	assert.Equal(t, strparse.ParseInput("bc"), i)
	assert.Equal(t, 'a', o)
}

func TestRuneOnMuitiBytes(t *testing.T) {
	p := strparse.Rune('あ')
	i, o, err := p.Parse("あいう")
	assert.NoError(t, err)
	assert.Equal(t, strparse.ParseInput("いう"), i)
	assert.Equal(t, 'あ', o)
}
