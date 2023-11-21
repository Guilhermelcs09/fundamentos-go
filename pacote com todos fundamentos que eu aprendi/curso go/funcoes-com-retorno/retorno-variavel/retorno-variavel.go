package main

import "fmt"

func soma(numero ...int) int {
	valor := 1
	for _, teste := range numero {
		valor += teste
	}
	return valor
}

func escrever(texto string, numero ...int) {
	for _, teste := range numero {
		fmt.Println(texto, teste)
	}
}

func main() {
	total := soma(1, 2)
	fmt.Println(total)

	escrever("ola", 10, 20, 30)
}
