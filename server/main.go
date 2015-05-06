package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	directory := flag.String("directory", "/usr/local/rct", "Path to the folder storing RCT experiment data")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: server [-directory DIRECTORY] ")
	}
	http.ListenAndServe(":8080", http.FileServer(http.Dir(*directory)))
}
