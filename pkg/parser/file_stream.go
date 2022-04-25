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

package parser

import (
	"io"
)

// IOReadSeeker is the stream input.
// We should require io.Seeker because some parsers try to recover from failures.
type IOReadSeeker struct {
	r io.ReadSeeker
}

// NewIOReadSeeker initializes an IOReadSeeker.
func NewIOReadSeeker(r io.ReadSeeker) *IOReadSeeker {
	return &IOReadSeeker{r}
}

// Read implements ParseInput interface.
func (r *IOReadSeeker) Read(buf []byte) (int, error) {
	return r.r.Read(buf)
}

// Seek implements ParseInput interface.
func (r *IOReadSeeker) Seek(n int, mode SeekMode) (int, error) {
	v, err := r.r.Seek(int64(n), int(mode))
	return int(v), err
}
