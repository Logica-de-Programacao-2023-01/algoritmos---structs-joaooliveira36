package main

import "fmt"

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func tempo_de_trabalho(funcionaio Funcionario) {
	tempo_trabalhando := funcionaio.idade - 18
	fmt.Printf("O tempo de trabalho de %s é de %d anos\n", funcionaio.nome, tempo_trabalhando)
}
func aumentar_salario(funcionario Funcionario) {
	salario_corrigido := 0.0
	if funcionario.salario <= 1000 {
		salario_corrigido = funcionario.salario + (funcionario.salario * 0.1)
	} else if funcionario.salario >= 2000 {
		salario_corrigido = funcionario.salario - (funcionario.salario * 0.1)
	} else if funcionario.salario < 2000 {
		salario_corrigido = funcionario.salario
	}

	fmt.Print("O salario corrigido é de ", salario_corrigido)
}

func main() {
	funcionario1 := Funcionario{
		nome:    "Rodrigo",
		salario: 3000,
		idade:   34,
	}
	tempo_de_trabalho(funcionario1)
	aumentar_salario(funcionario1)
}
