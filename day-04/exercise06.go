/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos
Demonstre estes valores
*/

package main

import (
	"fmt"
)

const (
	_ = 2024 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Println("---Anos---")
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d", d)
}
