package genetic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetOldFiles() []string {
	fmt.Println(*directory)
	// for d in directory/experiments:
	//	load the meta.json
	//	if the experiment is older than 1hour:
	//		find the iterations directory
	//		print it
	expDir := filepath.Join(*directory, "experiments")
	finfo, err := ioutil.ReadDir(expDir)
	if err != nil {
		panic(err)
	}
	for _, dir := range finfo {
		metaPath := filepath.Join(expDir, dir.Name(), "meta.json")
		em, err := readMetadata(metaPath)
		if err != nil {
			panic(err)
		}
		fmt.Println(em.Hash)
	}
	return []string{}
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
