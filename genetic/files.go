package genetic

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func GetOldFiles(d time.Duration) []string {
	expDir := filepath.Join(*directory, "experiments")
	finfo, err := ioutil.ReadDir(expDir)
	if err != nil {
		panic(err)
	}
	oldFolders := make([]string, 0)
	now := time.Now().UTC()
	for _, dir := range finfo {
		metaPath := filepath.Join(expDir, dir.Name(), "meta.json")
		em, err := readMetadata(metaPath)
		if err != nil {
			panic(err)
		}
		duration := now.Sub(em.Date)
		if duration > d {
			oldFolders = append(oldFolders, filepath.Join(expDir, dir.Name()))
		}
	}
	return oldFolders
}

// encode writes the given interface to the disk.
func encode(name string, v interface{}) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	return enc.Encode(v)
}

func readMetadata(filename string) (ExperimentMetadata, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var em ExperimentMetadata
	js := json.NewDecoder(f)
	err = js.Decode(&em)
	return em, err
}
