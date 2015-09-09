package main

import "testing"
import "github.com/bmizerany/assert"

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
