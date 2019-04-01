#!/bin/bash

# This file is targeted at development and manual testing needs.
# To run in production, build an executable and use that.
# This file can serve as a guideline for building. See below.

pushd $GOPATH/src/github.com/exercism/exalysis >/dev/null

# building starts here
GO111MODULE=on go generate ./...
GO111MODULE=on go build -o ./exalysis ./cmd/exalysis
# building ends here. Move exalysis executable to a folder contained in PATH variable.

popd >/dev/null
mv $GOPATH/src/github.com/exercism/exalysis/exalysis .
./exalysis
