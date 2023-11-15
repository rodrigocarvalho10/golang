/*
Crie um programa que:
- Atribua um valor int a uma variavel
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa varáivel 1 para a esquerda e atribua o resultado a outra váriavel
- Demonstre esta outra variável em decimal, binário e hexadecimal
*/

package main

import (
	"fmt"
)

var x = 500

func main() {
	fmt.Println("O Valor da variável é", x)
	fmt.Println("Decimal - Binário - Hexadecimal")
	fmt.Printf("%d, %b, %#x\n", x, x, x)
	y := x << 1
	fmt.Println("O Valor da variável é", y)
	fmt.Println("Decimal - Binário - Hexadecimal")
	fmt.Printf("%d, %b, %#x", y, y, y)
}
