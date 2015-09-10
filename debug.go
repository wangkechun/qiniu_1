package main

import (
	"qiniu_1/Godeps/_workspace/src/github.com/bmizerany/assert"
	"qiniu_1/Godeps/_workspace/src/github.com/k0kubun/pp"
	godebug "qiniu_1/Godeps/_workspace/src/github.com/tj/go-debug"
)

var print = pp.Print
var printf = pp.Printf
var println = pp.Println
var debug = godebug.Debug("single")
var equal = assert.Equal
