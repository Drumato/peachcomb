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

package byteparse

import (
	"encoding/binary"
	"unsafe"
)

var (
	envEndian binary.ByteOrder
)

func init() {
	determineEnvEndian()
}

func determineEnvEndian() {
	buf := [4]byte{}
	const value = uint32(0x01234567)
	bufPtr := unsafe.Pointer(&buf[0])
	// assignment
	*(*uint32)(bufPtr) = value

	switch buf {
	case [4]byte{0x67, 0x45, 0x23, 0x01}:
		envEndian = binary.LittleEndian
	case [4]byte{0x01, 0x23, 0x45, 0x67}:
		envEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}
