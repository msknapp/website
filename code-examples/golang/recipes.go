package golang

import (
	"io/ioutil"
)

func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func Reverse(s string) string {
	x := []byte(s)
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}
	return string(x)
}

func FileToString(filepath string) string {
	bits, e := ioutil.ReadFile(filepath)
	if e != nil {
		panic(e)
	}
	return string(bits)
}
