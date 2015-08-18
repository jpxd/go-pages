# go-pages

A wiki tool built on golang with git as the storage back-end. Content
is formatted in [markdown
syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet). The wiki is
rendered with go templates and [bootstrap](http://getbootstrap.com) css.
This project was forked from [aspic/g-wiki](https://github.com/aspic/g-wiki).

## Install

Simply go get it:

	go get github.com/jpxd/go-pages

then run it

	go-pages -http=:8080 -dir=files

## Develop

Templates are embedded with [go.rice](https://github.com/GeertJohan/go.rice).
If you change a file under templates or static, either use

	make debug

to execute go-pages and load the changed files live; or

	make build

to regenerate templates.rice-box.go and static.templates.rice-box.go, and create a portable binary.
