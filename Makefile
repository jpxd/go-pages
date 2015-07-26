run:
	rm -f templates.rice-box.go
	go build
	open http://localhost:8000/
	./g-wiki -http=:8000 -dir=files

build:
	go get github.com/GeertJohan/go.rice/rice
	rice embed-go
	go build
