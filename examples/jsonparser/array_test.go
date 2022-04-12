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

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSONArrayValue(t *testing.T) {
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonValueString("foo"),
			jsonValueString("bar"),
			jsonValueString("baz"),
		},
		length: 3,
	}
	i, o, err := parseJSONArrayValue([]rune(`["foo", "bar", "baz"]`))
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
	assert.Equal(t, "", string(i))
}

func TestParseJSONArrayValueWith2d(t *testing.T) {
	s := []rune(`[["a", "b"], ["c", "d"], ["e", "f"]]`)
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonArrayValue{
				elements: []jsonValue{
					jsonValueString("a"),
					jsonValueString("b"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonValueString("c"),
					jsonValueString("d"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonValueString("e"),
					jsonValueString("f"),
				},
				length: 2,
			},
		},
		length: 3,
	}
	i, o, err := parseJSONArrayValue(s)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
	assert.Equal(t, "", string(i))
}
