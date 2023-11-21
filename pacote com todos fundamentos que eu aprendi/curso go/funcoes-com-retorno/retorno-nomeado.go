package main

import "fmt"

func CaulculoMatematico(n1, n2 int16) (soma int16, sub int16) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := CaulculoMatematico(10, 10)
	fmt.Println(soma, sub)
}
