package main

import (
	"github.com/bmizerany/assert"
	"github.com/k0kubun/pp"
	godebug "github.com/tj/go-debug"
)

var print = pp.Print
var printf = pp.Printf
var println = pp.Println
var debug = godebug.Debug("bigNum")
var equal = assert.Equal
