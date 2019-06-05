package main

import "github.com/davecgh/go-spew/spew"

func inDumpWithSpew(data interface{}) {
	spew.Dump(data)
}
