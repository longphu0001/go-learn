#!/bin/bash

go test $(go list ./... | grep -v vendor | grep -v _examples) -coverprofile=coverage.out
# go test -coverprofile=coverage.out && go tool cover -html=coverage.out
# go test -race -v -coverprofile=coverage.txt -covermode=atomic ./

# PROJECT_FOLDER=`basename "$PWD"`

# for d in $(go list ./... | grep -v vendor | grep -v _examples); do

#     parentdir=`dirname "$d"`
#     subdir=`basename "$d"`
#     echo "======================="
#     echo $d
#     echo $parentdir
#     echo $subdir
#     # if [[ $subdir == "tests" ]]; then
#     #     parentdir="$parentdir/..."
#     #     go test -race -cover -coverpkg=$parentdir -coverprofile=profile.out $d
#     # else
#     #     go test -race -cover -coverprofile=profile.out $d
#     # fi
#     if [[ $parentdir != "." ]]; then
#         go test -race -cover -coverprofile=profile.out $d
#     fi

#     if [ -f profile.out ]; then
#         echo "Merging coverage file for" $d
#         tail -n+2 profile.out >>  "$PWD/coverage.out"
#         rm profile.out
#     fi

#     echo "======================="
# done

# go tool cover -html=coverage.out



