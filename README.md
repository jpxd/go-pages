# go-pages

A wiki tool built on golang with git as the storage back-end. Content
is formatted in [markdown
syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet). The wiki is
rendered with go templates and [bootstrap](http://getbootstrap.com) css but doesn't depend on any CDN.
This project was forked from [aspic/g-wiki](https://github.com/aspic/g-wiki).

## Using

Simply go get it:

	go get github.com/jpxd/go-pages

Initialize the git repo for the pages

	mkdir files
	cd files
	git init

Run it with

	go-pages

Usable command line flags are:

* -address=:8080
* -dir=files *(data directory has to be an intialized git repository!)*
* -title=CoolWiki

## Extensions

*TODO...*