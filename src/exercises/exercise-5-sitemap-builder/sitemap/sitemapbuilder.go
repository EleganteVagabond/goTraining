package sitemap

import (
	"encoding/xml"
	"exercises/exercise-4-link-parser/link"
	"io"
	"log"
	"net/http"
	"strings"
)

// siteList map structure, keys are hrefs, bool indicates we've visited before
var siteList map[string]bool

const xmlns = `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`

func init() {
	siteList = make(map[string]bool)
}

// WriteToXML writes the contents of the sitelist to a writer using xml semantics
func WriteToXML(writer io.Writer) {
	writer.Write([]byte(xml.Header))
	writer.Write([]byte(xmlns))
	EncodeListToXML(xml.NewEncoder(writer))
	writer.Write([]byte("</urlset>"))
}

// EncodeListToXML encodes the list stored in this object to an xml encoder
func EncodeListToXML(xmlEnc *xml.Encoder) {
	xmlEnc.Indent("\t", "\t")
	type url struct {
		Loc string `xml:"loc"`
	}
	for href, visited := range siteList {
		if visited {
			xmlEnc.Encode(url{href})
		}
	}
	xmlEnc.Flush()
}

// gets the full path for the url given a current base path (directory)
// if urlIn is an absolute path (i.e. http://something.com) then it is returned directly
// otherwise the basepath is prepended to the urlIn
func getFullPathURL(urlIn string, basePath string) string {
	ret := urlIn
	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}
	if strings.HasPrefix(urlIn, "/") {
		// this goes to the domain
		ret = basePath + urlIn[1:]
	}
	if strings.Index(ret, "://") == -1 {
		// doesn't have a transport in it
		ret = strings.TrimSpace(basePath + ret)
	}
	return ret
}

// check if the given url is in our domain
func isInDomain(url string, domain string) bool {
	return strings.Index(url, domain) != -1 || strings.HasPrefix(url, "/")
}

// buildSitemap builds a sitemap for a given domain
func buildSitemap(location string, siteRoot string, domain string, client http.Client) {
	if !siteList[location] && isInDomain(location, domain) {
		siteList[location] = true
		log.Println("getting links for page", location)
		resp, err := client.Get(location)
		if err != nil {
			log.Fatal(err)
		}
		loc, err := resp.Location()
		if loc != nil {
			// redirected, use this path, too
			siteList[getFullPathURL(loc.RawPath, domain)] = true
		}

		links := link.ParseHTMLLinks(resp.Body)
		defer resp.Body.Close()
		for _, link := range links {
			buildSitemap(getFullPathURL(link.Href, siteRoot), siteRoot, domain, client)
		}
	}
}

// PopulateSitemap populates the site map stored in this object
func PopulateSitemap(startLoc string) {
	startLoc, siteroot, domain := extractSiteInfo(startLoc)
	buildSitemap(startLoc, siteroot, domain, http.Client{})
}

// pulls out the site info (domain, root, startlocation) given a
// starting location (which might be google.com so needs more data)
func extractSiteInfo(startLoc string) (string, string, string) {
	var domain, siteroot string
	xportix := strings.Index(startLoc, "://")
	if xportix == -1 {
		// prepend transport
		startLoc = "http://" + startLoc
	}
	xportix = strings.Index(startLoc, "://")
	domain = startLoc[xportix+3:]
	slashIx := strings.Index(domain, "/")
	if slashIx != -1 {
		domain = domain[:slashIx]
	}
	siteroot = startLoc[:xportix+3] + domain

	return startLoc, siteroot, domain
}
