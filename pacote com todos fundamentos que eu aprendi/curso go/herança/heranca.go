package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"gui", "lucio", 19}
	p2 := estudante{p1, "veterinario", "unb"}
	fmt.Println(p2)

}
