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

package strparse_test

import (
	"strings"
	"testing"

	"github.com/Drumato/peachcomb/pkg/strparse"
	"github.com/stretchr/testify/assert"
)

func TestDigit1Failure(t *testing.T) {
	p := strparse.Digit1()

	i := strparse.NewCompleteInput("aabbccdd11223344")
	_, _, err := p(i)
	assert.Error(t, err)
}

// BenchmarkDigit1 benchmarks the performance of the Digit1 parser.
func BenchmarkDigit1(b *testing.B) {
	inputString := strings.Repeat("9", 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		input := strparse.NewCompleteInput(inputString)
		_, _, err := strparse.Digit1()(input)
		if err != nil {
			b.Fatalf("Failed to parse: %v", err)
		}
	}
}
