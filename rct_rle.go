// Dealing with RCT's run length encoding

// This should follow the API laid out by the gzip package.
// Author: Kevin Burke

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Reader struct {
	r   *bufio.Reader
	off int
}

type Writer struct {
	w      io.Writer
	closed bool
	err    error
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Create a new reader that can read encoded files.
func NewReader(r io.Reader) *Reader {
	z := new(Reader)
	z.r = bufio.NewReader(r)
	z.off = 0
	return z
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Compute the checksum for a given series of bytes.
func checksum(ride []byte) uint32 {
	summation := uint32(0)
	for i := range ride {
		tmp := summation + uint32(ride[i])
		summation = (summation & 0xffffff00) | (tmp & 0x000000ff)
		highthree := summation & (7 << 29)
		summation = summation << 3
		summation |= highthree >> 29
	}
	summation -= 120001
	return summation
}

// Read the encoded file into the byte array, decoding it in the process.
// Note, a TD6 file will have a checksum appended on the end. This checksum is
// not removed by the Read function.
func (z *Reader) Read(ride []byte) (int, error) {
	var n int
	if len(ride) == 0 {
		return 0, nil
	}
	// Read one byte
	c, err := z.r.ReadByte()
	if err != nil {
		return 0, err
	}
	// if positive, read that number of bytes
	if c >= 0 && c <= 128 {
		// copy c bytes from the reader to the byte array
		n, err = z.r.Read(ride[:min(int(c)+1, len(ride))])
		if err != nil {
			return 0, err
		}
		z.off += int(n)

	} else {
		// if negative, read the next byte, repeat that byte (-c + 1) times
		repeatByte, err := z.r.ReadByte()
		repeatTimes := 0 - int8(c) + 1
		rpts := bytes.Repeat([]byte{repeatByte}, int(repeatTimes))
		repeatReader := bytes.NewReader(rpts)
		n, err = repeatReader.Read(ride[:int(repeatTimes)])
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// implement the writing algorithm. Need to return the byte array so that we
// can compute checksums.
func (z *Writer) write(ride []byte) ([]byte, int, error) {
	if len(ride) == 0 {
		return []byte{}, 0, nil
	}
	if len(ride) == 1 {
		towrite := []byte{byte(1), ride[0]}
		n, err := z.w.Write(towrite)
		if err != nil {
			return []byte{}, n, err
		}
		return towrite, 1, nil
	}
	// if you find an immediate duplicate, read all of the duplicates to get
	// the count. write the count, then write the duplicate byte.
	if len(ride) > 1 && ride[0] == ride[1] {
		count := 2
		for {
			if len(ride) > count && ride[count-1] == ride[count] && count <= 127 {
				count += 1
			} else {
				break
			}
		}

		inv := -count + 1
		towrite := []byte{byte(inv), ride[0]}
		n, err := z.w.Write(towrite)
		if err != nil {
			return []byte{}, n, err
		}
		return towrite, count, nil
	} else {
		// read bytes until you find ones that are the same
		count := 1
		for {
			if len(ride)-1 == count {
				count += 1
				break
			}
			if ride[count] != ride[count+1] && count <= 127 {
				count += 1
			} else {
				break
			}
		}
		towrite := append([]byte{byte(len(ride[0:count]) - 1)}, ride[0:count]...)
		n, err := z.w.Write(towrite)
		if err != nil {
			return []byte{}, n, err
		}
		return towrite, count, nil
	}
}

func (z *Writer) Write(ride []byte) (int, error) {
	count := 0
	var b bytes.Buffer
	for {
		encoded, n, err := z.write(ride[count:])
		if err != nil {
			return n, err
		}
		b.Write(encoded)
		if count >= len(ride) {
			// all done. write checksum and continue
			cks := checksum(b.Bytes())
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, cks)
			if err != nil {
				return n, err
			}
			n, err := z.w.Write(buf.Bytes())
			if err != nil {
				return n, err
			}
			return count + n, nil
		}
		count += n
	}
}
