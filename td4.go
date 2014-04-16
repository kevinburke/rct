// td4 ride format parser

package main

import (
	//"encoding/hex"
	"fmt"
	"os"
)

func readRaptor() {
	f, err := os.Open("rides/raptor.td4")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 10)
	f.Read(b)
	fmt.Printf("%x", b)
}

func main() {
	f, err := os.Open("rides/hello.td4")
	fmt.Println(f)
	fmt.Println(err)
}
