#!/bin/bash

GOOS=linux GOARCH=386 go build ph && mv ph compiled/ph-linux-i686 && echo "linux i386"
GOOS=linux GOARCH=amd64 go build ph && mv ph compiled/ph-linux-x86_64 && echo "linux amd64"

GOOS=darwin GOARCH=386 go build ph && mv ph compiled/ph-darwin-i686 && echo "darwin i386"
GOOS=darwin GOARCH=amd64 go build ph && mv ph compiled/ph-darwin-x86_64 && echo "darwin amd64"
