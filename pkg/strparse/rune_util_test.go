package strparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPrefix(t *testing.T) {
	var s1, s2 []rune
	s1 = []rune("abcde")
	s2 = []rune("abc")
	assert.True(t, hasPrefix(s1, s2))

	s1 = []rune("abc")
	s2 = []rune("abc")
	assert.True(t, hasPrefix(s1, s2))

	s1 = []rune("abe")
	s2 = []rune("abc")
	assert.False(t, hasPrefix(s1, s2))

	s1 = []rune("")
	s2 = []rune("abc")
	assert.False(t, hasPrefix(s1, s2))
}
