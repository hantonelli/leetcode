package main

import (
	"math/big"
)

func Multiply2(num1 string, num2 string) string {
	n1, ok1 := new(big.Int).SetString(num1, 10)
	n2, ok2 := new(big.Int).SetString(num2, 10)
	if !ok1 || !ok2 {
		return ""
	}
	var res big.Int
	return res.Mul(n1, n2).String()
}
