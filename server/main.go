package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"path"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/kevinburke/rct"
	"github.com/kevinburke/rct/geo"
)

const VERSION = "0.1"

var directory string
var templateDirectory string
var staticDirectory string

func RenderError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func RenderTemplate(w http.ResponseWriter, name string, context interface{}) {
	funcMap := template.FuncMap{
		"truncate": func(l int, s string) string {
			var numRunes = 0
			for index, _ := range s {
				numRunes++
				if numRunes > l {
					return s[:index]
				}
			}
			return s
		},
	}

	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles(
		path.Join(templateDirectory, "base.html"),
		path.Join(templateDirectory, name),
	)
	if err != nil {
		RenderError(w, err)
		return
	}
	b := new(bytes.Buffer)
	err = tmpl.Execute(b, context)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, b)
	} else {
		RenderError(w, err)
	}
}

func MemberRenderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := LoadMember(vars["experiment"], vars["iteration"], vars["member"])
	if err != nil {
		RenderError(w, err)
		return
	}
	left, _ := geo.Render(m.Track)
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	err = enc.Encode(left)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, b)
	} else {
		RenderError(w, err)
	}
}

func MemberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	member, err := LoadMember(vars["experiment"], vars["iteration"], vars["member"])
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderTemplate(w, "member.html", member)
}

func IterationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iter, err := LoadIteration(vars["experiment"], vars["iteration"])
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderTemplate(w, "iteration.html", iter)
}

func ExperimentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exp, err := LoadExperiment(vars["experiment"])
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderTemplate(w, "experiment.html", exp)
}

func ExperimentListHandler(w http.ResponseWriter, r *http.Request) {
	experiments, err := LoadExperiments()
	if err != nil {
		RenderError(w, err)
		return
	}
	data := struct {
		Experiments []Experiment
	}{
		experiments,
	}
	RenderTemplate(w, "index.html", data)
}

func serverHeaderHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fmt.Sprintf("rct/%s", VERSION))
		for {
			rint := rand.Intn(len(rct.RIDENAMES))
			if rct.RIDENAMES[rint] == "(none)" {
				continue
			}
			w.Header().Set("X-Powered-By", rct.RIDENAMES[rint])
			break
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.StringVar(&directory, "directory", "/usr/local/rct",
		"Path to the folder storing RCT experiment data")
	flag.StringVar(&templateDirectory, "template-directory", "/usr/local/rct/server/templates",
		"Path to the folder storing RCT server templates (this file -> server/templates)")
	flag.StringVar(&staticDirectory, "static-directory", "/usr/local/rct/server/static",
		"Path to the folder storing RCT server templates (this file -> server/static)")

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: server [-directory DIRECTORY] ")
	}

	r := mux.NewRouter()
	r.HandleFunc("/experiments/{experiment}/iterations/{iteration}/{member}/render",
		MemberRenderHandler)
	r.HandleFunc("/experiments/{experiment}/iterations/{iteration}/{member}",
		MemberHandler)
	r.HandleFunc("/experiments/{experiment}/iterations/{iteration}",
		IterationHandler)
	r.HandleFunc("/experiments/{experiment}",
		ExperimentHandler)
	r.HandleFunc("/",
		ExperimentListHandler)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(staticDirectory))))
	http.Handle("/",
		serverHeaderHandler(r))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
