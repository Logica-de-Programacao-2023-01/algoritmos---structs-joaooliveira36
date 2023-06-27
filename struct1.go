package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func area_circ(circulo Circulo) float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func main() {
	raio := Circulo{
		raio: 10,
	}
	area := area_circ(raio)

	fmt.Print(area)
}
