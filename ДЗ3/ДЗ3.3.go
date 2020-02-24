package main

import (
	"github.com\stack"
	"fmt"
)

func main() {
stack.Push("Этот текст")
stack.Push("Будет находитьься в стеке")
stack.Push("До первого обращения в pop")

fmt.Println(stack.Pop())
fmt.Println(stack.Pop())

stack.Push("Добавим еще текста")

fmt.Println(stack.Pop())
fmt.Println(stack.Pop())

}