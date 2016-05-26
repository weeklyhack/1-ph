#!/bin/bash

GOOS=linux GOARCH=386 go build ph && mv ph compiled/ph-linux-i386 && echo "linux i386"
GOOS=linux GOARCH=amd64 go build ph && mv ph compiled/ph-linux-amd64 && echo "linux amd64"

GOOS=darwin GOARCH=386 go build ph && mv ph compiled/ph-darwin-i386 && echo "darwin i386"
GOOS=darwin GOARCH=amd64 go build ph && mv ph compiled/ph-darwin-amd64 && echo "darwin amd64"
