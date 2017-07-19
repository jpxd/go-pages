/*
GNU GPLv3 - see LICENSE
*/

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

// RevisionLogLimit limits the maximum amount of revisions shown for a page
const RevisionLogLimit = 5

var (
	// default values, can be overriden by flags
	directory           = "files"
	address             = ":8080"
	title               = "gopages"
	executableDirectory = "."
)

func init() {
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDirectory = path.Dir(executablePath)
}

func main() {
	// Define command line flags and parse them
	flagDirectory := flag.String("dir", directory, "directory where the markdown files are stored")
	flagAddress := flag.String("address", address, "address for the webserver to bind to, example: 0.0.0.0:8000")
	flagTitle := flag.String("title", title, "title to display")
	flag.Parse()

	// Update global variables to possibly overriden ones
	directory = *flagDirectory
	address = *flagAddress
	title = *flagTitle

	// Check if wiki data directory exists
	if _, err := os.Stat(directory); err != nil {
		log.Fatalln("WARNING: the specified directory (%q) does not exist!", directory)
	}

	// Static files (js, css, etc)
	staticDirectory := path.Join(executableDirectory, "static")
	fileServer := http.FileServer(http.Dir(staticDirectory))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Wiki handlers
	http.HandleFunc("/", wikiHandler)

	// Listen
	log.Printf("Start listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
