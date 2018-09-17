#!/bin/bash

pushd $GOPATH/src/github.com/tehsphinx/exalysis >/dev/null
GO111MODULE=on go build -o ./exalysis ./cmd
popd >/dev/null
mv $GOPATH/src/github.com/tehsphinx/exalysis/exalysis .
./exalysis
