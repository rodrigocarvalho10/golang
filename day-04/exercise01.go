/*
Escreva um programa que mostre um número em decimal, binário e hexadecimal
*/

package main

import (
	"fmt"
)

func main() {
	x := 50
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x", x)
}
