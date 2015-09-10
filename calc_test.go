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

func TestString(t *testing.T) {
	var x bigNum
	in := "12,54,45.12,6,8,7"
	x.input(in, 60)
	assert.Equal(t, x.String(), in)
}

func TestCalc(t *testing.T) {
	var x, y bigNum
	x.input("12,54,45.12,6,8,7", 60)
	y.input("1,1,0,1.0,1,0,1,1", 2)
	// pp.Println("x", x)
	// pp.Println("y", y)
}

func TestChangeBase(t *testing.T) {
	var x bigNum
	x.input("1,2,0,2,2", 3)
	// pp.Println(x)
	x.changeBaseInteger(6)
	// pp.Println('@', x)
	// pp.Println(x.String())
	assert.Equal(t, x.String(), "3,5,5")
}
func TestChangeBase2(t *testing.T) {
	var x bigNum
	x.input("0.1,2,0,2", 3)
	// pp.Println(x)
	x.changeBaseDecimal(6)
	// pp.Println('@', x)
	// pp.Println(x.String())
	assert.Equal(t, x.String(), "0.3,2,5,2")
}
