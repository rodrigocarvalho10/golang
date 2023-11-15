/*
Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
° ==
° !=
° <=
° <
° >=
° >
*/

package main

import (
	"fmt"
)

func main() {
	a := (9 == 9)
	b := (9 != 9)
	c := (9 <= 9)
	d := (9 < 9)
	e := (9 >= 9)
	f := (9 > 9)

	fmt.Printf("Valor de A é: %v\n", a)
	fmt.Printf("Valor de B é: %v\n", b)
	fmt.Printf("Valor de C é: %v\n", c)
	fmt.Printf("Valor de D é: %v\n", d)
	fmt.Printf("Valor de E é: %v\n", e)
	fmt.Printf("Valor de F é: %v", f)

}
