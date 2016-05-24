package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/kevinburke/rct/genetic"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	packageRoot := flag.String("package-root", "/", "Path to the folder running github.com/kevinburke/rct (for finding commit hashes)")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("Usage: genetic [-directory DIRECTORY] ")
	}
	log.Fatal(genetic.Run(*packageRoot))
}
