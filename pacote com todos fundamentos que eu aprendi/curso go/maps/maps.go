package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"pessoa1": {
			"nome":   "gui",
			"cidade": "df",
		},
	}
	fmt.Println(usuario2)
}
