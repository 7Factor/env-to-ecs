#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

go test -v 7factor.io/_unittests
