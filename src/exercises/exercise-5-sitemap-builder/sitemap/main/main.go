package main

import (
	"exercises/exercise-5-sitemap-builder/sitemap"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	startFlag := flag.String("root", "", "Give the root location. If not prepended with the transport type http:// is assumed")
	flag.Parse()

	startLoc := *startFlag
	if startLoc == "" {
		log.Fatal("requires location to start")
	}

	sitemap.PopulateSitemap(startLoc)
	bb := strings.Builder{}
	sitemap.WriteToXML(&bb)
	fmt.Println(bb.String())
}
