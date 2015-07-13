// Decode a TD6 file and print the results to stdout. It's going to be a blob
// of binary data, you probably want to redirect it to a file.
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
	// xxx is there a better way to do this?
	rr := io.TeeReader(reader, os.Stdout)
	var bitbuffer bytes.Buffer
	bitbuffer.ReadFrom(rr)
}
