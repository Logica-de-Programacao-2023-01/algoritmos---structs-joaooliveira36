package main

import "fmt"

type Endereço struct {
	rua    string
	numero string
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    string
	endereço Endereço
}

func Endereço_completo(pessoa1 Pessoa) string {
	return pessoa1.nome + " tem " + pessoa1.idade + " anos, " + "mora na rua " + pessoa1.endereço.rua +
		" numero " + pessoa1.endereço.numero + " em " + pessoa1.endereço.cidade + pessoa1.endereço.estado
}

func main() {
	pessoa1 := Pessoa{
		nome:  "Joao",
		idade: "18",
		endereço: Endereço{
			rua:    "Aos",
			numero: "5",
			cidade: "Brasília",
			estado: "DF",
		},
	}
	fmt.Print(Endereço_completo(pessoa1))
}
