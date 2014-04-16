package main

import (
	"encoding/hex"
	"io/ioutil"
)

func main() {
	b := "0047FF6F0564206A6F6221"
	bits, err := hex.DecodeString(b)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("rides/hello.td4", bits, 0644)
}
