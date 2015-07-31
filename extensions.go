package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"path"
)

type WikiExtension func(opts map[string]string, path string) (md string, err error)

var extensions map[string]WikiExtension

func registerExtension(key string, processFunction WikiExtension) {
	extensions[key] = processFunction
}

func optionFallback(m map[string]string, key string, fallback string) string {
	if value, ok := m[key]; ok {
		return value
	}
	return fallback
}

func parseExpression(exp string) (name string, options map[string]string, err error) {
	exp = strings.Trim(exp, "{}")
	exp = strings.TrimSpace(exp)
	parts := strings.Split(exp, " ")

	if len(parts) < 1 {
		return "", nil, errors.New("empty expression")
	}

	name = parts[0]
	parts = parts[1:]

	options = make(map[string]string)
	for _, part := range parts {
		if len(part) < 2 {
			continue
		}
		kv := strings.Split(part, ":")
		options[kv[0]] = "true"
		if len(kv) > 1 {
			options[kv[0]] = kv[1]
		}
	}

	return name, options, nil
}

func executeExpression(exp, path string) (res string, err error) {
	name, options, err := parseExpression(exp)
	if err != nil {
		return "", err
	}
	extension, ok := extensions[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("unknown extension '%s' in %s", name, path))
	}
	return extension(options, path)
}

func processExtensions(node *Node) {
	if len(extensions) == 0 {
		return
	}

	hay := node.Bytes
	buffer := bytes.Buffer{}
	startDelimiter, endDelimiter := []byte("{{"), []byte("}}")

	for {
		expStart := bytes.Index(hay, startDelimiter)
		if expStart < 0 {
			break
		}
		buffer.Write(hay[:expStart])
		hay = hay[expStart:]
		expEnd := bytes.Index(hay, endDelimiter)
		if expEnd < 0 {
			break
		}
		expEnd += len(endDelimiter)

		exp := string(hay[:expEnd])

		result, err := executeExpression(exp, node.Path)
		if err != nil {
			buffer.Write(hay[:expEnd])
			log.Println(err)
		} else {
			buffer.WriteString(result)
		}
		hay = hay[expEnd:]

	}
	buffer.Write(hay)
	node.Bytes = buffer.Bytes()
}

func registerAllExtensions() {
	registerExtension("tree", treeExtension)
}

func treeExtension(opts map[string]string, url string) (md string, err error) {
	dirpath := path.Join(directory, url)
	files, err := ioutil.ReadDir(dirpath)
	if len(files) == 0 {
		return "Such emptyness...", nil
	}
	buffer := bytes.Buffer{}
	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), ".md")
		if strings.HasPrefix(name, ".") || len(name) == 0 {
			continue
		}
		buffer.WriteString(" * [")
		buffer.WriteString(name)
		buffer.WriteString("](")
		buffer.WriteString(url)
		buffer.WriteString("/")
		buffer.WriteString(name)
		buffer.WriteString(")\n")
	}

	return buffer.String(), nil
}

func init() {
	extensions = make(map[string]WikiExtension)
	registerAllExtensions()
}
