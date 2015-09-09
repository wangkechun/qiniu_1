package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

var _ = pp.Println

func init() {
	log.SetFlags(log.Lshortfile)
}

type bigNum struct {
	integer []int
	decimal []int
	base    int
}

func (x *bigNum) input(s string, base int) {
	//l := len(s)
	x.base = base
	r := strings.Split(s, ".")
	var left, right string
	if len(r) == 2 {
		left = r[0]
		right = r[1]
	} else if len(r) == 1 {
		left = r[0]
		right = ""
	} else {
		panic("num illegal")
	}
	x.integer = splitNum(left)

	reverse(x.integer)
	x.decimal = splitNum(right)
}

func splitNum(s string) (r []int) {
	nums := strings.Split(s, ",")
	for i := 0; i < len(nums); i++ {
		num, _ := strconv.Atoi(nums[i])
		r = append(r, num)
	}
	return
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
