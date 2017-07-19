package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"strings"
)

func treeExtension(params []string, url string) (string, error) {
	url = url[:strings.LastIndex(url, "/")]
	if len(params) > 0 {
		url = path.Join(url, params[0])
	}
	dirpath := path.Join(directory, url)
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "No files or directories in tree", nil
	}

	buffer := bytes.Buffer{}
	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), ".md")
		if f.IsDir() {
			if name == ".git" {
				continue
			}
			targetURL := strings.TrimSuffix(url, "/") + "/" + name
			buffer.WriteString(" * [*" + name + "*](" + targetURL + ")\n")
			continue
		}
		if strings.HasPrefix(name, ".") || len(name) == 0 {
			continue
		}
		targetURL := strings.TrimSuffix(url, "/") + "/" + name
		buffer.WriteString(" * [" + name + "](" + targetURL + ")\n")
	}

	return buffer.String(), nil
}
