build:
	find . -name ".DS_Store" -delete
	go get github.com/GeertJohan/go.rice/rice
	rice embed-go
	go build

debug:
	rm *.rice-box.go
	go build
	open http://localhost:8000/
	./g-wiki -http=:8000 -dir=files
