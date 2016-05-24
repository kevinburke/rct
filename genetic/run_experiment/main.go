package main

import (
	"flag"
	"log"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	packageRoot := flag.String("package-root", "/", "Path to the folder running github.com/kevinburke/rct (for finding commit hashes)")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("Usage: genetic [-directory DIRECTORY] ")
	}
	log.Fatal(genetic.Run(*packageRoot))
}
