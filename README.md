# go-pages

A wiki tool built on golang with git as the storage back-end. Content is formatted in [commonmark syntax](https://spec.commonmark.org/0.30/). The wiki is rendered with go templates, [bootstrap](http://getbootstrap.com) css and [highlightjs](https://highlightjs.org) for code highlighting but doesn't depend on any CDN. This project was forked from [aspic/g-wiki](https://github.com/aspic/g-wiki).

## Using

Available command line flags are:

* `--address=:8080` *(in the format ip:port, empty ip binds to all ips)*
* `--dir=files` *(data directory has to be an intialized git repository!)*
* `--title=CoolWiki` *(title for the wiki)*
* `--basepath=/wiki/` *(base path for reverse proxy web applications)*

## Extensions

The goldmark rendering engine supports extensions which can be found here:

* https://github.com/yuin/goldmark/#built-in-extensions
* https://github.com/yuin/goldmark/#extensions

## Example screenshot

![Screenshot](static/screenshots/screenshot1.jpg)
