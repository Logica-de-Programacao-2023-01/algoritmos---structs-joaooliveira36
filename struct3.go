package main

import "fmt"

type Triângulo struct {
	base   float64
	altura float64
}

func area_triângulo(triangulo Triângulo) float64 {
	return (triangulo.base * triangulo.altura) / 2
}

func main() {
	triangulo1 := Triângulo{
		base:   10,
		altura: 8,
	}
	fmt.Print(area_triângulo(triangulo1))
}
