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

package byteparse_test

import (
	"testing"

	"github.com/Drumato/goparsecomb/pkg/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestUInt8(t *testing.T) {
	elfMagicNumber := [4]byte{0x7f, 0x45, 0x4c, 0x46}
	p := byteparse.UInt8()
	i, o, err := p.Parse(elfMagicNumber[:])
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x45, 0x4c, 0x46}, i)
	assert.Equal(t, uint8(0x7f), o)
}
