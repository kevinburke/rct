// Prints the latest track
package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/kevinburke/rct/genetic"
)

func main() {
	files, err := genetic.GetExperimentFiles()
	if err != nil {
		panic(err)
	}
	mostRecentFile := ""
	mostRecentDate := time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC)
	for _, dir := range files {
		metaPath := filepath.Join(dir, "meta.json")
		em, err := genetic.ReadMetadata(metaPath)
		if err != nil {
			panic(err)
		}
		if em.Date.Sub(mostRecentDate) > 0 {
			mostRecentDate = em.Date
			mostRecentFile = dir
		}
	}
	fmt.Println(mostRecentDate)
	fmt.Println(mostRecentFile)
}
