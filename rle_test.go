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
	z, err := NewReader(bitreader)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(z)
	out := bitbuffer.Bytes()
	if string(out) != "Good job!" {
		t.Fatalf("expected \"Good job!\" but got %s", out)
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
