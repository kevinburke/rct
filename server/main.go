package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"pkg/text/template"
	"regexp"

	"github.com/gorilla/handlers"
	"github.com/kevinburke/rct"
	"github.com/kevinburke/rct/genetic"
)

const VERSION = "0.1"

func renderHandler(directory string, templateDirectory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{status:\"ok\"}"))
	}
}

func newRctHandler(directory string, templateDirectory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := path.Join(directory, fmt.Sprintf("%s.json", r.URL.Path))
		f, err := os.Open(p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		dec := json.NewDecoder(f)
		m := new(genetic.Member)
		err = dec.Decode(m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

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

		tmpl, err := template.New("iteration.html").Funcs(funcMap).ParseFiles(path.Join(templateDirectory, "iteration.html"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		err = tmpl.Execute(w, m)
		if err != nil {
			w.Write([]byte("\n" + err.Error()))
		}
	}
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
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
	directory := flag.String("directory", "/usr/local/rct", "Path to the folder storing RCT experiment data")
	templateDirectory := flag.String("template-directory", "/usr/local/rct/server/templates", "Path to the folder storing RCT server templates (this file -> server/templates)")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: server [-directory DIRECTORY] ")
	}
	h := new(RegexpHandler)
	renderRoute, err := regexp.Compile("/experiments/.*/iterations/.*/.*/render")
	if err != nil {
		log.Fatal(err)
	}
	iterationRoute, err := regexp.Compile("/experiments/.*/iterations/.*/.*")
	if err != nil {
		log.Fatal(err)
	}
	defaultRoute, err := regexp.Compile("/")
	if err != nil {
		log.Fatal(err)
	}
	h.HandleFunc(renderRoute, renderHandler(*directory, *templateDirectory))
	h.HandleFunc(iterationRoute, newRctHandler(*directory, *templateDirectory))
	h.Handler(defaultRoute, http.FileServer(http.Dir(*directory)))
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, serverHeaderHandler(h))))
}
