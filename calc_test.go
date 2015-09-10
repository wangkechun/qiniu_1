package main

import (
	"testing"
)

func TestInput(t *testing.T) {
	x := New("[12,54,45.12,6,8,7]", 60)
	r := BigNum{
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
	equal(t, x, r)
}

func TestString(t *testing.T) {
	in := "[12,54,45.12,6,8,7]"
	x := New(in, 60)
	equal(t, x.String(), in)
}

func TestChangeBase(t *testing.T) {
	x := New("[1,2,0,2,2]", 3)
	x.changeBaseInteger(6)
	equal(t, x.RawString(), "[3,5,5]")
	equal(t, x.String(), "355")
}

func TestChangeBase2(t *testing.T) {
	x := New("[0.1,2,0,2]", 3)
	x.changeBaseDecimal(6)
	equal(t, x.RawString(), "[0.3,2,5,2]")
}

func TestShortInput(t *testing.T) {
	x := New("[1,1,0,1.0,1,0,1,1]", 2)
	y := New("1101.01011", 2)
	equal(t, x, y)
}

func TestShortInput2(t *testing.T) {
	x := New("abC012.c", 16)
	y := New("[10,11,12,0,1,2.12]", 16)
	equal(t, x, y)
}

func TestCalc(t *testing.T) {
	x := New("[12,54,45.12,6,8,7]", 60)
	y := New("[1,1,0,1.0,1,0,1,1]", 2)
	z := x.Add(&y)
	equal(t, z.String(), "[12,54,58.32,43,38,7]")
}

func TestCalc2(t *testing.T) {
	x := New("12", 3)
	y := New("11", 5)
	z := x.Add(&y)
	z.ChangeBase(10)
	equal(t, z.String(), "11")
}
