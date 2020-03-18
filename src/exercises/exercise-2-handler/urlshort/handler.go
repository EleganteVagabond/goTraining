package urlshort

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		if url != "" {
			http.Redirect(res, req, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
}

// DBHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the DB) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func DBHandler(fallback http.Handler) http.HandlerFunc {
	pathsToUrls, err := LoadURLMapFromDB()
	if err != nil {
		log.Fatal("Could not load DB Data")
	}
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		if url != "" {
			http.Redirect(res, req, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// use the facilities of yaml.v2 in order to parse the yaml data
	var ymlData []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yml, &ymlData); err != nil {
		return nil, err
	}
	// now build a map
	pathsToUrls := make(map[string]string, len(ymlData))
	for _, element := range ymlData {
		pathsToUrls[element.Path] = element.URL
	}
	// now create a functional handler that utilizes this map for speed and consistency
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := fmt.Sprintf("%v", pathsToUrls[req.URL.Path])
		if url != "" {
			http.Redirect(res, req, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonText []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse the JSON data
	var jsonData []struct {
		Path string
		URL  string
	}
	if err := json.Unmarshal(jsonText, &jsonData); err != nil {
		return nil, err
	}
	// now build a map
	pathsToUrls := make(map[string]string, len(jsonData))
	for _, element := range jsonData {
		pathsToUrls[element.Path] = element.URL
	}
	// now create a functional handler that utilizes this map for speed and consistency
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := fmt.Sprintf("%v", pathsToUrls[req.URL.Path])
		if url != "" {
			http.Redirect(res, req, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}), nil
}
