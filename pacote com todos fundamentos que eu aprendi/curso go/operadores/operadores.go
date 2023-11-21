package main

import "fmt"

func main() {

	var numero int16
	fmt.Print("digite um numero ")
	fmt.Scan(&numero)

	dobro := numero * 2

	fmt.Print("O dobro de ", numero, " é ", dobro)
}
