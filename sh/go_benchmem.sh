#!/bin/bash

# go test -v --bench . --benchmem
go test -coverprofile=coverage.out
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
# go test -race -v -coverprofile=coverage.txt -covermode=atomic ./




