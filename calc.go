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

func (x *bigNum) String() (r string) {
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

func splitNum(s string) (r []int) {
	if s == "" {
		return
	}
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

// func calcX(base, newBase, w int) {
// 	if newBase%base != 0 && newBase > base {
// 		panic("新base应该大于旧base并且是其倍数")
// 	}
// 	v := make([]int, w)
//
// }

func (x *bigNum) changeBaseInteger(newBase int) {
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
	if sum[length] != 0 {
		length++
	}
	x.integer = sum[:length]
}

func (x *bigNum) changeBaseDecimal(newBase int) {
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

	for length > 0 && length <= len(sum) && sum[length-1] == 0 {
		length--
	}
	x.decimal = sum[:length]
}

func (x *bigNum) changeBase(newBase int) {
	x.changeBaseInteger(newBase)
	x.changeBaseDecimal(newBase)
	x.base = newBase
}
func (x *bigNum) add(y *bigNum) (z bigNum) {
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
	//-----
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func trimRightZero(v []int) []int {
	i := len(v)
	for {
		if i > 0 && v[i-1] > 0 {
			break
		}
		i--
	}
	return v[:i]
}
