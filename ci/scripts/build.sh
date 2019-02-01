#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

go install 7factor.io/...