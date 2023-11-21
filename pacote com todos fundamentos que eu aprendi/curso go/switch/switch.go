package main

import "fmt"

func dia(numero int16) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça-feira"
	case 4:
		return "quarta"
	default:
		return "numero invalido"
	}

}

func main() {
	teste := dia(7)
	fmt.Println(teste)

}
