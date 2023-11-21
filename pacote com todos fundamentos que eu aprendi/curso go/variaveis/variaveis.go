package main

import "fmt"

func main() {
	var variavel1 string = "Guilherme"
	variavel2 := "flamenguista"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste"
		variavel4 string = "teste2"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "fla", "mengo"
	fmt.Println(variavel5, variavel6)
}
