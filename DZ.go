package main

import (
	"fmt"

)

func main() {
	var chislo int
	fmt.Println("Введите число")
	fmt.Scan(&chislo)
	if chislo%2 == 0 {
		fmt.Println(chislo, "число четное")
	} else {
		fmt.Println(chislo, "число не четное")
	}
}
