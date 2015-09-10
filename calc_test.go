package bignumadd

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/k0kubun/pp"
)

var _ = pp.Println

func TestInput(t *testing.T) {
	x := New("12,54,45.12,6,8,7", 60)
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
	assert.Equal(t, x, r)
}

func TestString(t *testing.T) {
	in := "12,54,45.12,6,8,7"
	x := New(in, 60)
	assert.Equal(t, x.String(), in)
}

func TestChangeBase(t *testing.T) {
	x := New("1,2,0,2,2", 3)
	x.changeBaseInteger(6)
	assert.Equal(t, x.String(), "3,5,5")
}

func TestChangeBase2(t *testing.T) {
	x := New("0.1,2,0,2", 3)
	x.changeBaseDecimal(6)
	assert.Equal(t, x.String(), "0.3,2,5,2")
}

func TestCalc(t *testing.T) {
	x := New("12,54,45.12,6,8,7", 60)
	y := New("1,1,0,1.0,1,0,1,1", 2)
	z := x.Add(&y)
	assert.Equal(t, z.String(), "12,54,58.32,43,38,7")
}
