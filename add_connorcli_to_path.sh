#!/bin/bash

go version # is go installed

go clean # remove any go binaries that may have been created before in this directory

go build -o ./connorcli # build connorcli

sudo mv connorcli /usr/local/bin/connorcli # Move it into the binaries so it can run in shell