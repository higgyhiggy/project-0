package main

import (
	"flag"
)

var file bool

func init() {
	flag.BoolVar(&file, "file", false, "file name to open or create")
	flag.Parse()
}
