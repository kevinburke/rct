package genetic

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// GetExperimentFiles returns a list of directories containing experiments
func GetExperimentFiles() (folders []string, err error) {
	expDir := filepath.Join(*directory, "experiments")
	finfo, err := ioutil.ReadDir(expDir)
	if err != nil {
		return folders, err
	}
	for _, dir := range finfo {
		path := filepath.Join(expDir, dir.Name())
		folders = append(folders, path)
	}
	return folders, nil
}

func GetOldFiles(d time.Duration) []string {
	var oldFolders []string
	files, err := GetExperimentFiles()
	if err != nil {
		panic(err)
	}
	now := time.Now().UTC()
	for _, dir := range files {
		metaPath := filepath.Join(dir, "meta.json")
		em, err := ReadMetadata(metaPath)
		if err != nil {
			panic(err)
		}
		duration := now.Sub(em.Date)
		if duration > d {
			oldFolders = append(oldFolders, dir)
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

func ReadMetadata(filename string) (ExperimentMetadata, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var em ExperimentMetadata
	js := json.NewDecoder(f)
	err = js.Decode(&em)
	return em, err
}
