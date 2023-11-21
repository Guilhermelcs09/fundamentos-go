package main

import "fmt"

func arrumar() {

	if r := recover(); r != nil {
		fmt.Println("Recuperado ")
	}
}

func aluno(n1, n2 float32) bool {
	defer arrumar()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Fodeu")
}

func main() {
	fmt.Println(aluno(6, 6))
}
