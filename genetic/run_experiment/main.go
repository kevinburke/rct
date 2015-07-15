package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	packageRoot := flag.String("package-root", "/", "Path to the folder running github.com/kevinburke/rct (for finding commit hashes)")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: genetic [-directory DIRECTORY] ")
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	log.Fatal(genetic.Run(*packageRoot))
}
