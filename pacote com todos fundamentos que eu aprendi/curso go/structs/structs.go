package main

import "fmt"

func main() {
	fmt.Println("Arquivo structs")

	type endereco struct {
		cidade string
		cep    int16
	}

	type registro struct {
		nome     string
		idade    int16
		endereco endereco
	}

	enderecoEx := endereco{"rua bobos", 15}

	usuario := registro{"gui", 19, enderecoEx}
	fmt.Println(usuario)

}
