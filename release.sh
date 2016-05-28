#!/bin/bash

GOOS=linux GOARCH=386 go build ph && mv ph compiled/ph-Linux-i686 && echo "linux i386"
GOOS=linux GOARCH=amd64 go build ph && mv ph compiled/ph-Linux-x86_64 && echo "linux amd64"

GOOS=darwin GOARCH=386 go build ph && mv ph compiled/ph-Darwin-i686 && echo "darwin i386"
GOOS=darwin GOARCH=amd64 go build ph && mv ph compiled/ph-Darwin-x86_64 && echo "darwin amd64"
