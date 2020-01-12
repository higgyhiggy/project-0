package main

import (
	"flag"
)

var name bool

func init() {
	flag.BoolVar(&name, "file", false, "file name to open or create")
	flag.Parse()
}
