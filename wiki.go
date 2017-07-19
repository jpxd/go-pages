package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

var baseTemplate = template.New("wiki")

func init() {
	// Load base templates for reusing
	_, err := baseTemplate.ParseFiles("templates/header.tpl", "templates/footer.tpl",
		"templates/edit.tpl", "templates/revisions.tpl",
		"templates/revision.tpl", "templates/node.tpl")
	if err != nil {
		log.Fatal(err)
	}
}

// Node holds a Wiki node.
type Node struct {
	Title    string
	Path     string
	File     string
	Content  string
	Template string
	Revision string
	Bytes    []byte
	Dirs     []*Directory
	Log      []*Log
	Markdown template.HTML

	Edit      bool // Edit mode
	Revisions bool // Show revisions
	Author    string
	Changelog string
}

// Directory lists nodes.
type Directory struct {
	Path   string
	Name   string
	Active bool
}

// Log is an event in the past.
type Log struct {
	Hash    string
	Message string
	Time    string
	Link    bool
}

func (node *Node) isHead() bool {
	return len(node.Log) > 0 && node.Revision == node.Log[0].Hash
}

// ToMarkdown processes the node contents.
func (node *Node) ToMarkdown() {
	node.ProcessExtensions()
	node.Markdown = template.HTML(string(blackfriday.MarkdownCommon(node.Bytes)))
}

func wikiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	// Params
	content := r.FormValue("content")
	changelog := r.FormValue("msg")
	author := r.FormValue("author")
	reset := r.FormValue("revert")
	revision := r.FormValue("revision")

	// Default to index page on trailing slash
	if r.URL.Path[len(r.URL.Path)-1] == '/' {
		r.URL.Path += "index"
	}
	filePath := fmt.Sprintf("%s%s.md", directory, r.URL.Path)
	node := &Node{
		File:  r.URL.Path[1:] + ".md",
		Path:  r.URL.Path,
		Title: title,
	}
	node.Revisions = parseBool(r.FormValue("revisions"))
	node.Edit = parseBool(r.FormValue("edit"))

	if cookie, err := r.Cookie("author"); err == nil {
		node.Author = cookie.Value
	}
	if node.Author == "" {
		node.Author = "Unknown"
	}

	node.Dirs = listDirectories(r.URL.Path)

	// We have content, update
	if content != "" && changelog != "" && author != "" {
		node.Author = author
		bytes := []byte(content)
		err := writeFile(bytes, filePath)
		if err != nil {
			log.Printf("Cant write to file %q, error: %v", filePath, err)
		} else {
			// Wrote file, commit
			node.Bytes = bytes
			node.GitAdd().GitCommit(changelog, author).GitLog()
			node.ToMarkdown()
		}
	} else if reset != "" {
		// Reset to revision
		node.Revision = reset
		node.GitRevert().GitCommit("Reverted to: "+node.Revision, author)
		node.Revision = ""
		node.GitShow().GitLog()
		node.ToMarkdown()
	} else {
		// Show specific revision
		node.Revision = revision
		node.GitShow().GitLog()

		createNew := len(node.Bytes) == 0
		node.Edit = node.Edit || createNew

		changelogPageName := strings.TrimLeft(node.Path, "/")
		if changelogPageName == "" {
			changelogPageName = "index page"
		}
		node.Changelog = fmt.Sprintf("Edit %s", changelogPageName)
		if createNew {
			node.Changelog = fmt.Sprintf("Create %s", changelogPageName)
		}

		if node.Edit {
			node.Content = string(node.Bytes)
			node.Template = "edit.tpl"
		} else {
			node.ToMarkdown()
		}
	}
	renderTemplate(w, node)
}

func writeFile(bytes []byte, entry string) error {
	err := os.MkdirAll(path.Dir(entry), 0777)
	if err == nil {
		return ioutil.WriteFile(entry, bytes, 0644)
	}
	return err
}

func setCookie(w http.ResponseWriter, name, value string) {
	expiration := time.Now().AddDate(1, 0, 0)
	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func renderTemplate(w http.ResponseWriter, node *Node) {
	// Set cookies
	setCookie(w, "author", node.Author)

	// Clone base template
	t, err := baseTemplate.Clone()
	if err != nil {
		log.Fatalln("Could not clone baseTemplate:", err)
	}

	// Build content template
	if node.Markdown != "" {
		tpl := "{{ template \"header\" . }}"

		// Show revisions
		if node.Revisions {
			tpl += "{{ template \"revisions\" . }}"
		}

		if !node.isHead() && node.Revision != "" {
			tpl += "{{ template \"revision\" . }}"
		}
		// Add node
		tpl += "{{ template \"node\" . }}"

		// Footer
		tpl += "{{ template \"footer\" . }}"
		if t, err = t.Parse(tpl); err != nil {
			log.Fatalf("Couldn't parse template %q: %v", tpl, err)
		}
		// Execute
		err = t.Execute(w, node)
	} else if node.Template != "" {
		err = t.ExecuteTemplate(w, node.Template, node)
	}
	if err != nil {
		log.Fatal("Could not execute template: ", err)
	}

}
