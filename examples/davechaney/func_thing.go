package math

import "github.com/pnsafonov/penny/generic"

type ThisNumberType generic.Number

func ThisNumberTypeMax(fn func(a, b ThisNumberType) bool, a, b ThisNumberType) ThisNumberType {
	if fn(a, b) {
		return a
	}
	return b
}
