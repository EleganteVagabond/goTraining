package urlshort

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testURL  = "/test"
	testDest = "http://localhost/unit-test-handler"

	bogusURL = "http://www.bogus.com"

	fallbackBody    = "fallback"
	fallbackHandler = http.HandlerFunc(fallback)
)

type testArgs struct {
	url         string
	destination string
	body        string
}

// function signature to pass to handler function for handling the "fallback" request
func fallback(w http.ResponseWriter, r *http.Request) {
	// Fprint is used to print to a writer, as opposed to stdout or similar
	fmt.Fprint(w, fallbackBody)
}

func TestMapHandler(t *testing.T) {
	// construct our test map. We assume if one works any will work
	// could add additional test cases for unicode, ascii type compatibility issues on different platforms
	pathsToUrls := map[string]string{testURL: testDest}
	// construct our tester
	mpTester := MapHandler(pathsToUrls, fallbackHandler)
	doCommonHandlerTests(mpTester, nil, t)
}

func TestYAMLHandler(t *testing.T) {
	// testYaml
	// make sure that if we pass malformed yaml it doesn't work
	{
		_, err := YAMLHandler([]byte("bogus"), fallbackHandler)
		if err == nil {
			t.Error("does not throw error for malformed YAML")
		}
	}
	yaml := []byte(`
- path: ` + testURL + `
  url: ` + testDest + `
`)
	yamlTester, err := YAMLHandler(yaml, fallbackHandler)
	doCommonHandlerTests(yamlTester, err, t)
}

func TestJSONHandler(t *testing.T) {
	// testJSON
	// make sure that if we pass malformed yaml it doesn't work
	{
		_, err := JSONHandler([]byte("bogus"), fallbackHandler)
		if err == nil {
			t.Error("does not throw error for malformed JSON")
		}
	}
	json := []byte(`[
 { "path": "` + testURL + `","url": "` + testDest + `" }
	]`)
	jsonTester, err := JSONHandler(json, fallbackHandler)
	doCommonHandlerTests(jsonTester, err, t)
}

func doCommonHandlerTests(handler http.HandlerFunc, err error, t *testing.T) {
	t.Helper()
	//make sure there are no errors
	checkNil(err, t)
	// test that the url is redirected
	validateHandlerTest(handler, t, testArgs{url: testURL, destination: testDest, body: ""})
	// test that the url is NOT redirected and we get the fallback body from the fallback handler
	validateHandlerTest(handler, t, testArgs{url: bogusURL, destination: "", body: fallbackBody})
}

func validateHandlerTest(mpTester http.HandlerFunc, t *testing.T, args testArgs) {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, args.url, nil)
	response := httptest.NewRecorder()
	mpTester(response, request)
	result := response.Result()

	if args.destination != "" {
		url, err := result.Location()
		checkNil(err, t)
		if url.String() != args.destination {
			t.Error("URL", url, "should be ", args.destination)
		}
	}

	if args.body != "" {
		// test the body is the same
		body, err := ioutil.ReadAll(result.Body)
		checkNil(err, t)
		if string(body) != args.body {
			t.Error("Body", body, "should be ", args.body)
		}
	}
}

func checkNil(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
