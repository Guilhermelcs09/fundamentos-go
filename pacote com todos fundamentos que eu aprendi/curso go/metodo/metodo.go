package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salva() {
	fmt.Println("salvando dados")
}

func (u *usuario) acrescentar() {
	u.idade++
}

func main() {
	usuario1 := usuario{"guilherme", 19}
	usuario1.salva()
	usuario1.acrescentar()
	fmt.Println(usuario1)
}
