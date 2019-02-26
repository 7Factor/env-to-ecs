#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin
export GOCACHE=off

go test -v 7factor.io/_inttests
