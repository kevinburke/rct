package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/rct/genetic"
)

type Experiment struct {
	ID         string
	Path       string
	Iterations []Iteration
}

func (e Experiment) URL() string {
	return filepath.Join("experiments", e.ID)
}

type Iteration struct {
	Number  string
	URL     string
	Members []Member
}

type Member struct {
	ID  string
	URL string
}

func LoadMember(exp, number, member string) (*genetic.Member, error) {
	p := filepath.Join(directory, "experiments", exp, "iterations", number, member+".json")
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	m := new(genetic.Member)
	err = dec.Decode(m)
	return m, err
}

func LoadIteration(exp, number string) (Iteration, error) {
	iteration := Iteration{URL: "", Number: number}
	prefix := filepath.Join(directory, "experiments", exp, "iterations", number)
	files, err := filepath.Glob(filepath.Join(prefix, "*"))
	if err != nil {
		return iteration, err
	}
	members := make([]Member, len(files))
	for i, path := range files {
		members[i] = Member{
			URL: strings.Replace(strings.Replace(path, directory, "", -1), ".json", "", -1),
			ID:  strings.Replace(strings.Replace(path, prefix+"/", "", -1), ".json", "", -1),
		}
	}
	iteration.Members = members
	return iteration, nil
}

func LoadExperiments() ([]Experiment, error) {
	prefix := filepath.Join(directory, "experiments")
	files, err := filepath.Glob(filepath.Join(prefix, "*"))
	if err != nil {
		return nil, err
	}
	experiments := make([]Experiment, len(files))
	for i, path := range files {
		experiments[i] = Experiment{
			Path: path,
			ID:   strings.Replace(path, prefix+"/", "", -1),
		}
	}
	return experiments, nil
}

func LoadExperiment(name string) (Experiment, error) {
	exp := Experiment{}
	prefix := filepath.Join(directory, "experiments", name, "iterations")
	files, err := filepath.Glob(filepath.Join(prefix, "*"))
	if err != nil {
		return exp, err
	}
	iterations := make([]Iteration, len(files))
	for i, path := range files {
		iterations[i] = Iteration{
			URL:    strings.Replace(path, directory, "", -1),
			Number: strings.Replace(path, prefix+"/", "", -1),
		}
	}
	exp.Iterations = iterations
	return exp, nil
}
