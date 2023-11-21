package main

import "fmt"

func main() {

	var variavel = 100
	var local = &variavel
	fmt.Println(variavel, local)

	variavel = 20
	fmt.Println(variavel, local)
}
