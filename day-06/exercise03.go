package main

import "fmt"

func main() {
	dtNasc := 1992
	anoAtual := 2024

	for dtNasc <= anoAtual {
		fmt.Println(dtNasc)
		dtNasc++
	}

}
