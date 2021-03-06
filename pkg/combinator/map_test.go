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
	"fmt"
	"testing"

	"github.com/Drumato/peachcomb/pkg/combinator"
	"github.com/Drumato/peachcomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestMapSubParserFailure(t *testing.T) {
	subP := strparse.Rune('a')
	p := combinator.Map(subP, func(ch rune) (bool, error) { return ch == 'a', nil })

	i := strparse.NewCompleteInput("bc")
	_, _, err := p(i)
	assert.Error(t, err)
}

func TestMapFnFailure(t *testing.T) {
	subP := strparse.Rune('a')
	p := combinator.Map(subP, func(ch rune) (bool, error) { return false, fmt.Errorf("") })

	i := strparse.NewCompleteInput("a")
	_, _, err := p(i)
	assert.Error(t, err)
}
