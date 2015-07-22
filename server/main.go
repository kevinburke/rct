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
	"github.com/kevinburke/rct/rle"
	"github.com/kevinburke/rct/td6"
	"github.com/kevinburke/rct/tracks"
)

const VERSION = "0.2"

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

func td6Handler(directory string, templateDirectory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		iterPath := strings.Replace(r.URL.Path, "/td6", "", 1)
		m, err := loadMember(directory, iterPath)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		query := r.URL.Query()
		completeParam := query.Get("complete")
		complete := completeParam == "true"
		// For whatever reason RCT2 needs the extension to be in uppercase
		shortName := fmt.Sprintf("%s.TD6", truncate(17, m.Id))
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", shortName))
		w.Header().Set("Content-Type", "text/td6")
		ride := td6.CreateMineTrainRide(m.Track, complete)
		for i := 0; i < len(ride.TrackData.Elements); i++ {
			fmt.Println(tracks.ElementNames[ride.TrackData.Elements[i].Segment.Type])
		}
		bits, err := td6.Marshal(ride)
		paddedBits := td6.Pad(bits)
		rleWriter := rle.NewWriter(w)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		rleWriter.Write(paddedBits)
	}
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

type tmplData struct {
	Member *genetic.Member
	Path   string
}

func truncate(l int, s string) string {
	var numRunes = 0
	for index, _ := range s {
		numRunes++
		if numRunes > l {
			return s[:index]
		}
	}
	return s
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
			"truncate": truncate,
		}

		tmpl, err := template.New("iteration.html").Funcs(funcMap).ParseFiles(path.Join(templateDirectory, "iteration.html"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		b := new(bytes.Buffer)
		err = tmpl.Execute(b, tmplData{Member: m, Path: r.URL.Path})
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

func buildRoute(regex string) *regexp.Regexp {
	route, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	return route
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
	td6Route := buildRoute("/experiments/.+/iterations/.+/.+/td6")
	renderRoute := buildRoute("/experiments/.*/iterations/.+/.+/render")
	iterationRoute := buildRoute("/experiments/.+/iterations/.+/.+")
	staticRoute := buildRoute("/static")
	defaultRoute := buildRoute("/")

	h.HandleFunc(td6Route, td6Handler(*directory, *templateDirectory))
	h.HandleFunc(renderRoute, renderHandler(*directory, *templateDirectory))
	h.HandleFunc(iterationRoute, newRctHandler(*directory, *templateDirectory))
	h.Handler(staticRoute, http.StripPrefix("/static", http.FileServer(http.Dir(*staticDirectory))))
	h.Handler(defaultRoute, http.FileServer(http.Dir(*directory)))
	allHandlers := jsonStripper(serverHeaderHandler(h))
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, allHandlers)))
}
