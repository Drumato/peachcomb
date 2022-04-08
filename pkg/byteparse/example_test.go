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
	"encoding/binary"
	"fmt"

	"github.com/Drumato/goparsecomb/pkg/byteparse"
)

func ExampleUInt8() {
	i, o, err := byteparse.UInt8().Parse([]byte{0x01, 0x02, 0x03})
	fmt.Println(i)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// [2 3]
	// 1
	// <nil>
}

func ExampleUInt16() {
	i, o, err := byteparse.UInt16(binary.BigEndian).Parse([]byte{0x01, 0x02, 0x03})
	fmt.Println(i)
	fmt.Printf("0x%x\n", o)
	fmt.Println(err)
	// Output:
	//
	// [3]
	// 0x102
	// <nil>
}