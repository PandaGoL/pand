package main

import (
	"fmt"
)

func main() {
	var chislo int
	fmt.Println("Введите число")
	fmt.Scan(&chislo)
	if chislo%3 == 0 {
		fmt.Println(chislo, "делится на 3 без остатка")
	} else {
		fmt.Println(chislo, "не делится на 3 без остатка")
	}
}
