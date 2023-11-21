package main

import "fmt"

func nota(n1, n2 float32) bool {
	defer fmt.Println("Nota calculada")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(nota(10, 16))
}
