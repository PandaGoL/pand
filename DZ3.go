package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(2)
	
	for i := 0; i < 100; i++{ 
		if i == 0 {
			fmt.Println("0")
		} else	if i != 0 && i <= 2 {
			fmt.Println("1")
	} else {
		a = a.Add(a, b)
		a, b = b, a
		fmt.Println(a)
	}
	}
}
