package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/kevinburke/rct/rle"
)

func main() {
	filename := flag.String("filename", "", "td6 file to decode")
	flag.Parse()
	if *filename == "" {
		log.Fatal("please enter a filename")
	}
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := rle.NewReader(f)
	rr := io.TeeReader(reader, os.Stdout)
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(rr)
}
