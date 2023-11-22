package main

import "fmt"

func main() {

	feriado := "sim"
	switch feriado {
	case "não":
		fmt.Println("Hoje não é feriado, vai trabalhar normal")
	case "sim":
		fmt.Println("Feriado!!! - Bora descansar")
	}
}
