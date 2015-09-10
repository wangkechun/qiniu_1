package main

import (
	"strconv"
	"strings"
)

// BigNum 大数存储
type BigNum struct {
	// 整数部分， 0位置代表个位
	integer []int
	// 小数部分， 0代表小数点后面第一位
	decimal []int
	// 该数字进制
	base int
}

const numTable = "0123456789abcdefghijklmnopqrstuvwxyz"

// Input 输入一个大数
func (x *BigNum) Input(s string, base int) {
	x.base = base
	charType := true
	if s[0] == '[' && s[len(s)-1] == ']' {
		s = s[1 : len(s)-1]
		charType = false
	}

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
	if charType {
		x.integer = splitChar(left)
		x.decimal = splitChar(right)
	} else {
		x.integer = splitNum(left)
		x.decimal = splitNum(right)
	}
	reverse(x.integer)
	x.integer = trimRightZero(x.integer)
	x.decimal = trimRightZero(x.decimal)
	x.Format()
}

// RawString 转换成 [1,2,3] 类似的形式
func (x *BigNum) RawString() (r string) {
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
	r = "[" + r + "]"
	return
}

func (x *BigNum) String() (r string) {
	if x.base >= 10+26 {
		return x.RawString()
	}
	var s []byte
	for i := len(x.integer) - 1; i >= 0; i-- {
		v := x.integer[i]
		s = append(s, numTable[v])
	}
	if len(x.decimal) > 0 {
		s = append(s, ',')
		for i := 0; i < len(x.decimal); i++ {
			v := x.decimal[i]
			s = append(s, numTable[v])
		}
	}
	return string(s)
}

func (x *BigNum) changeBaseInteger(newBase int) {
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
	if len(x.decimal) == 0 {
		return
	}
	if newBase%x.base != 0 || newBase < x.base {
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

// ChangeBase 修改数字的进制
func (x *BigNum) ChangeBase(newBase int) {
	x.changeBaseInteger(newBase)
	x.changeBaseDecimal(newBase)
	x.base = newBase
	x.Format()
}

// Add 大数相加
func (x *BigNum) Add(y *BigNum) (z BigNum) {
	newBase := x.base * y.base / gcd(x.base, y.base)
	x.ChangeBase(newBase)
	// log.Println("newBase", newBase)
	y.ChangeBase(newBase)
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
	z.Format()
	return
}

// New returns a new BigNum
func New(s string, base int) (v BigNum) {
	v.Input(s, base)
	return
}

//Format 格式化大数， 去掉多余的0
func (x *BigNum) Format() {
	x.integer = trimRightZero(x.integer)
	x.decimal = trimRightZero(x.decimal)
	if len(x.integer) == 0 {
		x.integer = []int{0}
	}
}
