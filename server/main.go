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
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/kevinburke/rct"
	"github.com/kevinburke/rct/genetic"
	"github.com/kevinburke/rct/geo"
)

const VERSION = "0.1"

func loadMember(directory string, pth string) (*genetic.Member, error) {
	p := path.Join(directory, fmt.Sprintf("%s.json", pth))
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	m := new(genetic.Member)
	err = dec.Decode(m)
	return m, err
}

func renderHandler(directory string, templateDirectory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		iterPath := strings.Replace(r.URL.Path, "/render", "", 1)
		m, err := loadMember(directory, iterPath)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}

func newRctHandler(directory string, templateDirectory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := loadMember(directory, r.URL.Path)
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
		b := new(bytes.Buffer)
		err = tmpl.Execute(b, m)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			io.Copy(w, b)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
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

func jsonStripper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			newUrl := strings.TrimSuffix(r.URL.Path, ".json")
			http.Redirect(w, r, newUrl, http.StatusMovedPermanently)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	directory := flag.String("server-directory", "/usr/local/rct", "Path to the folder storing RCT experiment data")
	templateDirectory := flag.String("template-directory", "/usr/local/rct/server/templates", "Path to the folder storing RCT server templates (this file -> server/templates)")
	staticDirectory := flag.String("static-directory", "/usr/local/rct/server/static", "Path to the folder storing RCT server templates (this file -> server/static)")

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: server [-directory DIRECTORY] ")
	}
	h := new(RegexpHandler)
	renderRoute, err := regexp.Compile("/experiments/.*/iterations/.+/.+/render")
	if err != nil {
		log.Fatal(err)
	}
	iterationRoute, err := regexp.Compile("/experiments/.*/iterations/.+/.+")
	if err != nil {
		log.Fatal(err)
	}
	staticRoute, err := regexp.Compile("/static")
	if err != nil {
		log.Fatal(err)
	}
	defaultRoute, err := regexp.Compile("/")
	if err != nil {
		log.Fatal(err)
	}
	h.HandleFunc(renderRoute, renderHandler(*directory, *templateDirectory))
	h.HandleFunc(iterationRoute, newRctHandler(*directory, *templateDirectory))
	h.Handler(staticRoute, http.StripPrefix("/static", http.FileServer(http.Dir(*staticDirectory))))
	h.Handler(defaultRoute, http.FileServer(http.Dir(*directory)))
	allHandlers := jsonStripper(serverHeaderHandler(h))
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, allHandlers)))
}
