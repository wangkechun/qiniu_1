package main

import "testing"
import "github.com/bmizerany/assert"

import "github.com/k0kubun/pp"

var _ = pp.Println

func TestInput(t *testing.T) {
	var x bigNum
	x.input("12,54,45.12,6,8,7", 60)
	r := bigNum{
		integer: []int{
			45,
			54,
			12,
		},
		decimal: []int{
			12,
			6,
			8,
			7,
		},
		base: 60,
	}
	assert.Equal(t, x, r)
}

func TestCalc(t *testing.T) {
	var x, y bigNum
	x.input("12,54,45.12,6,8,7", 60)
	y.input("1,1,0,1.0,1,0,1,1", 2)
	pp.Println("x", x)
	pp.Println("y", y)
}
