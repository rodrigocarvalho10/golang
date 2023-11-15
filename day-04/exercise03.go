/*
Crie constantes tipadas e não-tipadas
Demonstre seus valores
*/

package main

import (
	"fmt"
)

const (
	name     string = "Rodrigo"
	age      int    = 31
	children        = 2
)

func main() {
	fmt.Println(name, age, children)
}
