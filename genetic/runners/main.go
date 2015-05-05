package main

import (
	"flag"
	"log"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	directory := flag.String("directory", "/usr/local/rct", "Path to the folder storing RCT experiment data")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: genetic [-directory DIRECTORY] ")
	}
	log.Fatal(genetic.Run(*directory))
}
