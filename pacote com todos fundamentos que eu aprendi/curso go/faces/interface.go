package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma %0.2f", f.area())

}

type triangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}
