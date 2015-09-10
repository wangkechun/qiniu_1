package main

import (
	"github.com/bmizerany/assert"
	"github.com/k0kubun/pp"
	godebug "github.com/tj/go-debug"
)

var print = pp.Print
var debug = godebug.Debug("single")
var equal = assert.Equal
