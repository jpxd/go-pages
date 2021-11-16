FROM golang:1.16

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/go-pages
COPY . .

RUN go get -v
RUN go vet -v
RUN go install

RUN mkdir -p files
RUN git config --global init.defaultBranch main
RUN git config --global user.email "system@dockercontainer"
RUN git config --global user.name "system"
RUN git init files

EXPOSE 8080

ENTRYPOINT ["/go/bin/go-pages"]
