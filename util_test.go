package bignum

import "testing"

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3}
	reverse(s)
	r := []int{3, 2, 1}
	equal(t, s, r)
}

func TestSplitChar(t *testing.T) {
	r, _ := splitChar("asdgfsd123")
	equal(t, r, []int{10, 28, 13, 16, 15, 28, 13, 1, 2, 3})
}

func TestSplitNum(t *testing.T) {
	r, _ := splitNum("12,23,324,342")
	equal(t, r, []int{12, 23, 324, 342})
}

func TestSplitNum2(t *testing.T) {
	r, _ := splitNum("23,")
	equal(t, r, []int{23})
}

func TestSplitNum3(t *testing.T) {
	r, _ := splitNum(",12,23,")
	equal(t, r, []int{12, 23})
}

func TestSplitNum4(t *testing.T) {
	r, _ := splitNum(",")
	equal(t, len(r), 0)
}

func TestSplitNum5(t *testing.T) {
	r, _ := splitNum("")
	equal(t, len(r), 0)
}

func TestTrimRightZero(t *testing.T) {
	v := trimRightZero([]int{1, 2, 0, 0, 0})
	equal(t, v, []int{1, 2})
}

func TestTrimRightZero2(t *testing.T) {
	v := trimRightZero([]int{0, 0, 0})
	equal(t, v, []int{})
}

func TestGcd(t *testing.T) {
	equal(t, gcd(8, 12), 4)
}
