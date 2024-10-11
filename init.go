package main

import "flag"

var SilentMode bool

func init() {
	silent := flag.Bool("s", false, "Silent mode: don't open browser automatically")
	flag.BoolVar(silent, "silent", false, "Silent mode: don't open browser automatically")
	flag.Parse()
	SilentMode = *silent
}
