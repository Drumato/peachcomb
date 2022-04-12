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

	"github.com/Drumato/peachcomb/pkg/combinator"
	"github.com/Drumato/peachcomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestDelimitedBeginFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinator.Delimited(begin, contents, end)
	_, _, err := p([]rune("'12345\""))
	assert.Error(t, err)
}

func TestDelimitedContentsFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinator.Delimited(begin, contents, end)
	_, _, err := p([]rune("\"abcde\""))
	assert.Error(t, err)
}

func TestDelimitedEndFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinator.Delimited(begin, contents, end)
	_, _, err := p([]rune("\"12345'"))
	assert.Error(t, err)
}
