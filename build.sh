#!/bin/bash

for GOOS in darwin linux windows; do
  docker run --rm -v "$GOPATH":/go -v "$PWD":/usr/src/incrementor -w /usr/src/incrementor -e GOOS=$GOOS -e GOARCH=amd64 golang:1.8 go build -v -o incrementor-$GOOS
done
