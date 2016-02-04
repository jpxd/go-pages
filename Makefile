build:
	find . -name ".DS_Store" -delete
	go get github.com/GeertJohan/go.rice/rice
	rice embed-go
	go build

debug:
	rm -f rice-box.go
	go build
	mkdir files && cd files && git init
	./go-pages -http=:8000 -dir=files
