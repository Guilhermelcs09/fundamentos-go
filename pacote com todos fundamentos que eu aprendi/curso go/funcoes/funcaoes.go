package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func teste1(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	calculo := somar(10, 20)
	fmt.Println(calculo)

	var teste = func(n1 int8, n2 int8) int8 {
		fmt.Println(n1, n2)
		return n1 + n2
	}
	resultado := teste(10, 10)
	fmt.Println(resultado)

	resultadosoma, resultadosub := teste1(10, 5)
	fmt.Println(resultadosoma, resultadosub)
}
