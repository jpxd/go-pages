/*
GNU GPLv3 - see LICENSE
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// Register your extension methods here
var extensions = map[string]WikiExtension{
	"tree": treeExtension,
}

// WikiExtension is a function wich can be called from a page like {{ functionName param1 param2 ... }}
type WikiExtension func(params []string, path string) (markdown string, err error)

var extensionStartDelimiter, extensionEndDelimiter = []byte("{{"), []byte("}}")

func executeExpression(exp, path string) (markdown string, err error) {
	exp = strings.Trim(exp, "{}")
	parts := strings.Fields(exp)

	if len(parts) < 1 {
		return "", fmt.Errorf("Empty expression")
	}

	name := parts[0]
	extension, exists := extensions[name]
	if !exists {
		return "", fmt.Errorf("Unknown extension '%s' in %s", name, path)
	}

	return extension(parts[1:], path)
}

// ProcessExtensions processes all extension calls in the node contents and applies the content changes
func (node *Node) ProcessExtensions() {
	source := node.Bytes
	target := bytes.Buffer{}

	for {
		expStart := bytes.Index(source, extensionStartDelimiter)
		if expStart < 0 {
			break
		}
		target.Write(source[:expStart])
		source = source[expStart:]
		expEnd := bytes.Index(source, extensionEndDelimiter)
		if expEnd < 0 {
			target.Write([]byte("*Unmatched expression start*"))
			break
		}
		expEnd += len(extensionEndDelimiter)
		exp := string(source[:expEnd])

		result, err := executeExpression(exp, node.Path)
		if err != nil {
			target.Write(source[:expEnd])
			log.Println(err)
		} else {
			target.WriteString(result)
		}
		source = source[expEnd:]
	}
	target.Write(source)
	node.Bytes = target.Bytes()
}
