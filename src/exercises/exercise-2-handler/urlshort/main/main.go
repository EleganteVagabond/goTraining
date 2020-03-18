package main

import (
	"exercises/exercise-2-handler/urlshort"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var yamlFile string
	flag.StringVar(&yamlFile, "YAML", "", "Use to specify the location of a yaml file to load")
	var jsonFile string
	flag.StringVar(&jsonFile, "JSON", "", "Use to specify the location of a JSON file to load")
	var useDB bool
	flag.BoolVar(&useDB, "UseDB", false, "Flag for using a db (YAML/JSON flags ignored if true)")
	flag.Parse()

	var fallback http.Handler
	mux := defaultMux()

	// load data from db for handler
	if useDB {
		// create db handler using the mux as fallback
		fallback = urlshort.DBHandler(mux)
	} else {
		// use default
		fallback = mux
	}
	// Build the MapHandler using the mux or db as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, fallback)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	if yamlFile != "" {
		var err error
		yaml, err = ioutil.ReadFile(yamlFile)
		if err != nil {
			panic(err)
		}
	}
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the JSONHandler using the YAMLHandler as the
	// fallback
	json := `
  [ { "path": "/urlshortjson",
  "url": "https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/"},
  {"path": "/urlshortjson-final",
  "url": "https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go"
  }
  ]
`
	if jsonFile != "" {
		jsondata, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			panic(err)
		}
		json = string(jsondata)
	}
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
