package main

import (
	"fmt"
	"monkey/lexer"
)

func main() {
	hello := lexer.New("hello")
	fmt.Println(hello)
}
