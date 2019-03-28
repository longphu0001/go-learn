#!/bin/bash

go test -v --bench . --benchmem
go test -race -v -coverprofile=coverage.txt -covermode=atomic ./




