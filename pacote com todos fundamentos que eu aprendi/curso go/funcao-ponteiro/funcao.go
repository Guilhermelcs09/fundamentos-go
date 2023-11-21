package main

import "fmt"

func invertaSinal(numero int) int {
	return numero * -1

}

func main() {
	numero := 20
	numeroInvertido := invertaSinal(numero)
	fmt.Println(numeroInvertido)
}
