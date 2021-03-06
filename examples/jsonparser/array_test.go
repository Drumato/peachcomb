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

	"github.com/Drumato/peachcomb/pkg/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONArrayValue(t *testing.T) {
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonStringValue("foo"),
			jsonStringValue("bar"),
			jsonStringValue("baz"),
		},
		length: 3,
	}

	i := byteparse.NewCompleteInput([]byte(`["foo", "bar", "baz"]`))
	_, o, err := parseJSONArrayValue(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONArrayValueWith2d(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte(`[["a", "b"], ["c", "d"], ["e", "f"]]`))
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("a"),
					jsonStringValue("b"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("c"),
					jsonStringValue("d"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("e"),
					jsonStringValue("f"),
				},
				length: 2,
			},
		},
		length: 3,
	}
	_, o, err := parseJSONArrayValue(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}
