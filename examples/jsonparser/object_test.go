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

func TestParseJSONObjectValue_Empty(t *testing.T) {
	expected := jsonObjectValue{
		Fields: make([]jsonObjectField, 0),
	}
	i := byteparse.NewCompleteInput([]byte(`{ }`))
	_, o, err := parseJSONObjectValue(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectValue(t *testing.T) {
	expected := jsonObjectValue{
		Fields: []jsonObjectField{
			{
				Name:  "A",
				Value: jsonStringValue("B"),
			},
			{
				Name:  "C",
				Value: jsonStringValue("D"),
			},
		},
	}
	i := byteparse.NewCompleteInput([]byte(`{"A":"B","C":"D"}`))
	_, o, err := parseJSONObjectValue(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectField(t *testing.T) {
	expected := jsonObjectField{
		Name:  "foo",
		Value: jsonStringValue("bar"),
	}
	i := byteparse.NewCompleteInput([]byte(`"foo" : "bar"`))
	_, o, err := parseJSONObjectField(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectFieldList(t *testing.T) {
	expected := []jsonObjectField{
		{
			Name:  "n1",
			Value: jsonStringValue("v1"),
		},
		{
			Name:  "n2",
			Value: jsonStringValue("v2"),
		},
	}
	i := byteparse.NewCompleteInput([]byte(`"n1" : "v1", "n2" : "v2"`))
	_, o, err := parseJSONObjectFieldList(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}
