// MIT License
//
// Copyright (c) 2022 Drumato
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package combinator_test

import (
	"testing"

	"github.com/Drumato/goparsecomb/pkg/combinator"
	"github.com/Drumato/goparsecomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestTakewhile1(t *testing.T) {
	subP := strparse.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	p := combinator.TakeWhile1(subP)

	i, o, err := p.Parse([]rune("aaaabaa"))
	assert.NoError(t, err)
	assert.Equal(t, "aaaa", string(o))
	assert.Equal(t, "baa", string(i))
}

func TestTakewhile0(t *testing.T) {
	subP := strparse.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	p := combinator.TakeWhile0(subP)

	i, o, err := p.Parse([]rune("baa"))
	assert.NoError(t, err)
	assert.Equal(t, "", string(o))
	assert.Equal(t, "baa", string(i))
}