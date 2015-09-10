package main

import (
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
	i := len(v)
	for {
		if i > 0 && v[i-1] > 0 {
			break
		}
		i--
	}
	return v[:i]
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
