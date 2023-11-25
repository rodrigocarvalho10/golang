package main

import "fmt"

func main() {
	dtNasc := 1992
	anoAtual := 2024

	for {
		if dtNasc > anoAtual {
			break
		}
		fmt.Println(dtNasc)
		dtNasc++
	}

}
