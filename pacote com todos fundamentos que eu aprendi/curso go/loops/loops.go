package main

import (
	"fmt"
)

func main() {
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Adicionando")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)

	//	for j := 0; j < 5; j++ {
	//		fmt.Println("Adicionando", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"guilherme", "elaine", "lindo"}
	for posicao, produto := range nomes {
		fmt.Println(posicao, produto)
	}
}
