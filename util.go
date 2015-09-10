package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] > 0 {
			return v[:i+1]
		}
	}
	return []int{}
}

// 1,2,3
func splitNum(s string) (r []int) {
	if s == "" {
		return
	}
	nums := strings.Split(s, ",")

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if len(v) == 0 {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("illegal number %s", nums[i]))
		}
		r = append(r, num)
	}
	return
}

func splitChar(s string) (r []int) {
	for _, v := range s {
		n := int(v)
		if n >= '0' && n <= '9' {
			r = append(r, n-'0')
		} else if n >= 'a' && n <= 'z' {
			r = append(r, n-'a'+10)
		} else if n >= 'A' && n <= 'Z' {
			r = append(r, n-'A'+10)
		} else {
			panic(fmt.Sprintf("illegal character %v", v))
		}
	}
	return
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
