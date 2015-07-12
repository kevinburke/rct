package main

import (
	"fmt"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	files := genetic.GetOldFiles()
	for _, file := range files {
		fmt.Println(file)
	}
}
