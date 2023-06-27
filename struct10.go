package main

import (
	"fmt"
)

type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func (f Filme) AdicionarAvaliação(avaliação int) Filme {
	f.Avaliações = append(f.Avaliações, avaliação)
	return f
}

func (f Filme) RemoverAvaliação(posição int) Filme {
	if posição >= 0 && posição < len(f.Avaliações) {
		f.Avaliações = append(f.Avaliações[:posição], f.Avaliações[posição+1:]...)
	}
	return f
}

func (f Filme) CalcularMédiaAvaliações() float64 {
	total := 0
	for _, nota := range f.Avaliações {
		total += nota
	}
	return float64(total) / float64(len(f.Avaliações))
}

func (f Filme) ImprimirInformações() {
	fmt.Printf("Título: %s\n", f.Título)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Avaliações: %v\n", f.Avaliações)
	fmt.Printf("Média de Avaliações: %.2f\n", f.CalcularMédiaAvaliações())
}

func main() {
	filme := Filme{
		Título:  "Pulp Fiction",
		Diretor: "Quentin Tarantino",
		Ano:     1994,
		Avaliações: []int{
			5, 4, 4, 5, 3,
		},
	}

	filme = filme.AdicionarAvaliação(4)
	filme = filme.RemoverAvaliação(2)

	filme.ImprimirInformações()
}
