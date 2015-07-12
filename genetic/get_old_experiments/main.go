package main

import (
	"fmt"
	"time"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	files := genetic.GetOldFiles(1 * time.Hour)
	for _, file := range files {
		fmt.Println(file)
	}
}
