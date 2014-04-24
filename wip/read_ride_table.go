package main

import (
	"fmt"
	"os"
)

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/code/OpenRCT2/openrct2.exe")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 300)
	f.ReadAt(b, 0x0019889C)
	fmt.Println(b)
	for i := 0; i < 20; i++ {
		fmt.Printf("Ride %d\n", i)
		onebyte := int(b[i*4])
		if hasBit(onebyte, 1) {
			fmt.Println("has flat track")
		}
		if hasBit(onebyte, 2) {
			fmt.Println("has station platform")
		}
		if hasBit(onebyte, 6) {
			fmt.Println("has banking")
		}
		if hasBit(onebyte, 7) {
			fmt.Println("has vertical loop")
		}
		onebyte = int(b[i*4+1])
		if hasBit(onebyte, 0) {
			fmt.Println("has normal slope")
		}
		if hasBit(onebyte, 1) {
			fmt.Println("has steep slope")
		}
		if hasBit(onebyte, 2) {
			fmt.Println("has flat to steep")
		}
		fmt.Println("\n")
	}
}
