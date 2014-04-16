// Decompress a td4 ride file

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	//"os"
)

type Reader struct {
	r   *bufio.Reader
	off int
}

func NewReader(r io.Reader) (*Reader, error) {
	z := new(Reader)
	z.r = bufio.NewReader(r)
	z.off = 0
	return z, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (z *Reader) Read(ride []byte) (int, error) {
	var n int
	if len(ride) == 0 {
		return 0, nil
	}
	c, err := z.r.ReadByte()
	if err != nil {
		return 0, err
	}
	fmt.Printf("c is %d\n", c)
	if c >= 0 && c <= 128 {
		// copy c bytes from the reader to the byte array
		n, err = z.r.Read(ride[:min(int(c)+1, len(ride))])
		if err != nil {
			return 0, err
		}
		fmt.Println("read n bytes", n)
		z.off += int(n)
	} else {
		repeatByte, err := z.r.ReadByte()
		repeatTimes := 0 - int8(c) + 1
		fmt.Printf("repeating the number %d %d times\n", repeatByte, repeatTimes)
		fmt.Println(z.off)
		repeatReader := bytes.NewReader(
			bytes.Repeat([]byte{repeatByte}, int(repeatTimes)))
		n, err = repeatReader.Read(ride[:int(repeatTimes)])
		if err != nil {
			return 0, err
		}
		fmt.Println("read n bytes in repeat", n)
	}
	return n, nil
}
