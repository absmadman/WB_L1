package main

import (
	"fmt"
	"math/big"
)

func main() {
	ten := big.NewInt(10)
	a := big.NewInt(10)
	b := big.NewInt(10)
	for i := 0; i < 20; i++ {
		a.Mul(a, ten)
		b.Mul(b, ten)
	} // получили 2 числа больше чем 1 * 10^20
	res := big.NewInt(0)
	fmt.Println(res.Add(a, b)) // +
	fmt.Println(res.Sub(a, b)) // -
	fmt.Println(res.Mul(a, b)) // *
	fmt.Println(res.Div(a, b)) // /
}
