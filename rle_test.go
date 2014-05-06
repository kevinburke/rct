package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestDecompress(t *testing.T) {
	b := "0047FF6F0564206A6F6221"
	bits, err := hex.DecodeString(b)
	if err != nil {
		t.Fatalf(err.Error())
	}
	bitreader := bytes.NewReader(bits)
	z := NewReader(bitreader)
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	out := bitbuffer.Bytes()
	if string(out) != "Good job!" {
		t.Fatalf("expected \"Good job!\" but got %s", out)
	}
}

func TestCompress(t *testing.T) {
	var buf bytes.Buffer
	z := NewWriter(&buf)
	z.Write([]byte("Goood job!"))
	b := "0047fe6f0564206a6f6221d882451d"
	if hex.EncodeToString(buf.Bytes()) != b {
		t.Fatalf("expected %s but got %s", b, hex.EncodeToString(buf.Bytes()))
	}
}

//func TestMSB(t *testing.T) {
//thirtytwo := getmsb(32)
//if thirtytwo != 5 {
//t.Errorf("expected msb(32) = 5, was %d", thirtytwo)
//}

//thirtyone := getmsb(31)
//if thirtyone != 4 {
//t.Errorf("expected msb(31) = 4, was %d", thirtyone)
//}
//}
