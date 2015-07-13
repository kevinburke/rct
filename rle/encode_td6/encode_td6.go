package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/kevinburke/rct/rle"
)

func main() {
	usage := "Invalid arguments. Usage: encode_td6 filename"
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal(usage)
	}
	filename := args[0]
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	writer := rle.NewWriter(os.Stdout)
	writer.Write(bits)
}
