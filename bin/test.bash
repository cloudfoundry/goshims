#!/bin/bash

set -eu
set -o pipefail

#no tests in this repo, just make sure build passes
go build ./...
