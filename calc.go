package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

var p = pp.Println

func init() {
	log.SetFlags(log.Lshortfile)
}

// BigNum 大数字存储
type BigNum struct {
	// 整数部分， 0位置代表个位
	integer []int
	// 小数部分， 0代表小数点后面第一位
	decimal []int
	// 该数字进制
	base int
}

func (x *BigNum) input(s string, base int) {
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
		panic("number illegal" + s)
	}
	x.integer = splitNum(left)
	reverse(x.integer)
	x.decimal = splitNum(right)
}

func (x *BigNum) String() (r string) {
	var s []string
	for i := len(x.integer) - 1; i >= 0; i-- {
		s = append(s, strconv.Itoa(x.integer[i]))
	}
	r += strings.Join(s, ",")
	if len(x.decimal) > 0 {
		r += "."
		s = []string{}
		for i := 0; i < len(x.decimal); i++ {
			s = append(s, strconv.Itoa(x.decimal[i]))
		}
		r += strings.Join(s, ",")
	}
	return
}

// func calcX(base, newBase, w int) {
// 	if newBase%base != 0 && newBase > base {
// 		panic("新base应该大于旧base并且是其倍数")
// 	}
// 	v := make([]int, w)
//
// }

func (x *BigNum) changeBaseInteger(newBase int) {
	if newBase%x.base != 0 && newBase > x.base {
		panic("新base应该大于旧base并且是其倍数")
	}
	sum := make([]int, len(x.integer)+2)
	now := make([]int, len(x.integer)+2)
	length := 1
	now[0] = 1
	for i := 0; i < len(x.integer); i++ {
		for j := 0; j < length; j++ {
			sum[j] += x.integer[i] * now[j]
			sum[j+1] += sum[j] / newBase
			sum[j] %= newBase
			now[j] *= x.base
		}
		for j := 0; j < length; j++ {
			now[j+1] += now[j] / newBase
			now[j] %= newBase
		}
		if i != len(x.integer)-1 && now[length] != 0 {
			length++
		}
		// log.Printf("i=%d, length=%d\n", i, length)
		// log.Println("sum", sum)
		// log.Println("now", now)
		// log.Println()
	}
	x.integer = trimRightZero(sum)
}

func (x *BigNum) changeBaseDecimal(newBase int) {
	if newBase%x.base != 0 && newBase > x.base {
		panic("新base应该大于旧base并且是其倍数")
	}
	sum := make([]int, len(x.decimal)+2)
	now := make([]int, len(x.decimal)+2)
	length := 1
	now[0] = newBase / x.base
	for i := 0; i < len(x.decimal); i++ {
		for j := 0; j < length; j++ {
			sum[j] += x.decimal[i] * now[j]
			now[j] *= newBase / x.base
		}
		for j := length - 1; j > 0; j-- {
			sum[j-1] += sum[j] / newBase
			sum[j] %= newBase
		}
		for j := length - 1; j >= 0; j-- {
			now[j+1] = now[j] % newBase
			if j != 0 {
				now[j-1] += now[j] / newBase
			}
		}
		now[0] /= newBase

		if i != len(x.decimal)-1 && now[length] != 0 {
			length++
		}
		// log.Printf("i=%d, length=%d\n", i, length)
		// log.Println("sum", sum)
		// log.Println("now", now)
		// log.Println()
	}
	x.decimal = trimRightZero(sum)
}

func (x *BigNum) changeBase(newBase int) {
	x.changeBaseInteger(newBase)
	x.changeBaseDecimal(newBase)
	x.base = newBase
}
func (x *BigNum) add(y *BigNum) (z BigNum) {
	newBase := x.base * y.base / gcd(x.base, y.base)
	x.changeBase(newBase)
	// log.Println("newBase", newBase)
	y.changeBase(newBase)
	// log.Printf("x=%s  y=%s\n",x,y)
	integerLength := max(len(x.integer), len(y.integer))
	z.integer = make([]int, integerLength+2)
	for i := 0; i < integerLength; i++ {
		if i < len(x.integer) {
			z.integer[i] += x.integer[i]
		}
		if i < len(y.integer) {
			z.integer[i] += y.integer[i]
		}
		z.integer[i+1] += z.integer[i] / newBase
		z.integer[i] %= newBase
	}
	for z.integer[integerLength] != 0 {
		integerLength++
	}
	z.integer = z.integer[:integerLength+1]
	//decimal
	decimalLength := max(len(x.decimal), len(y.decimal))
	z.decimal = make([]int, decimalLength+2)
	for i := decimalLength - 1; i >= 0; i-- {
		if i < len(x.decimal) {
			z.decimal[i] += x.decimal[i]
		}
		if i < len(y.decimal) {
			z.decimal[i] += y.decimal[i]
		}
		if i != 0 {
			z.decimal[i-1] += z.decimal[i] / newBase
			z.decimal[i] %= newBase
		}
	}
	z.integer[0] += z.decimal[0] / newBase
	z.decimal[0] %= newBase
	z.base = newBase
	z.integer = trimRightZero(z.integer)
	z.decimal = trimRightZero(z.decimal)
	return
}

// New returns a new BigNum
func New(s string, base int) (v BigNum) {
	v.input(s, base)
	return
}
