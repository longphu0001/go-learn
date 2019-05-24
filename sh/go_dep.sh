#!/bin/bash

go get -v -u github.com/golang/dep/cmd/dep

dep status
dep ensure




