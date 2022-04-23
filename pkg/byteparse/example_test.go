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

	"github.com/Drumato/peachcomb/pkg/byteparse"
)

func ExampleUInt8() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03})
	_, o, err := byteparse.UInt8()(b)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// 1
	// <nil>
}

func ExampleUInt16() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03})
	_, o, err := byteparse.UInt16(binary.BigEndian)(b)
	fmt.Printf("0x%x\n", o)
	fmt.Println(err)
	// Output:
	//
	// 0x102
	// <nil>
}

func ExampleUInt32() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03, 0x04})
	_, o, err := byteparse.UInt32(binary.BigEndian)(b)
	fmt.Printf("0x%x\n", o)
	fmt.Println(err)
	// Output:
	//
	// 0x1020304
	// <nil>
}

func ExampleTag() {
	t := []byte{0x7f, 0x45, 0x4c, 0x46}

	b := byteparse.NewCompleteInput([]byte{0x7f, 0x45, 0x4c, 0x46, 0x02})
	_, o, err := byteparse.Tag(t)(b)
	fmt.Printf("%d\n", len(o))
	fmt.Printf("%x %x %x %x\n", o[0], o[1], o[2], o[3])
	fmt.Println(err)
	// Output:
	// 4
	// 7f 45 4c 46
	// <nil>
}
