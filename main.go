/*
GNU GPLv3 - see LICENSE
*/

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// RevisionLogLimit limits the maximum amount of revisions shown for a page
const RevisionLogLimit = 5

var (
	// default values, can be overriden by flags
	directory = "files"
	templates = "templates"
	static    = "static"
	address   = ":8080"
	title     = "gopages"
	basepath  = "/"
)

func main() {
	// Define command line flags and parse them
	flagDirectory := flag.String("dir", directory, "directory where the markdown files are stored")
	flagTemplates := flag.String("templates", templates, "directory where the templates are stored")
	flagStatic := flag.String("static", static, "directory where the static files are stored")
	flagAddress := flag.String("address", address, "address for the webserver to bind to, example: 0.0.0.0:8000")
	flagTitle := flag.String("title", title, "title to display")
	flagBasepath := flag.String("basepath", basepath, "base path, for web application proxy pass")
	flag.Parse()

	// Update global variables to possibly overriden ones
	directory = *flagDirectory
	templates = *flagTemplates
	static = *flagStatic
	address = *flagAddress
	title = *flagTitle
	basepath = *flagBasepath

	// Check if wiki data directory exists
	if _, err := os.Stat(directory); err != nil {
		log.Fatalf("WARNING: the specified directory (%q) does not exist!", directory)
	}

	err := loadTemplates(templates)
	if err != nil {
		log.Fatalf("Error loading templates from %s: %v", templates, err)
	}

	// Static files (js, css, etc)
	fileServer := http.FileServer(http.Dir(static))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Wiki handlers
	http.HandleFunc("/", wikiHandler)

	// Listen
	log.Printf("Start listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
