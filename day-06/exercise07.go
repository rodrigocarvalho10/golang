package main

import "fmt"

func main() {
	relacionamento := "namorando"

	if relacionamento == "casado" {
		fmt.Println("Ele é casado")
	} else if relacionamento == "namorando" {
		fmt.Println("Ele está namorando")
	} else {
		fmt.Println("Ele é solteiro")
	}
}
