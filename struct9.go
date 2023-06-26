package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}

	return soma / float64(len(a.Notas))
}

func (a *Aluno) ImprimirDados() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Média das notas:", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{7.5, 8.2, 6.9},
	}

	aluno.AdicionarNota(9.0)
	aluno.RemoverNota(1)

	aluno.ImprimirDados()
}
