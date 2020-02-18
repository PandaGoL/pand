package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(2)
	
	for i := 1; i < 100; i++ {
		a = a.Add(a, b)
		a, b = b, a
		fmt.Println(a)
	}
}
