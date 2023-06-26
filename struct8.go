package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preço   int
}

func ViagemMaisCara(viagens []Viagem) (Viagem, error) {
	if len(viagens) == 0 {
		return Viagem{}, fmt.Errorf("A lista de viagens está vazia")
	}
	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.preço > viagemMaisCara.preço {
			viagemMaisCara = viagem
		}
	}
	return viagemMaisCara, nil
}

func main() {
	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: "2023-07-01", preço: 250.0},
		{origem: "Rio de Janeiro", destino: "Salvador", data: "2023-08-15", preço: 350.0},
		{origem: "Salvador", destino: "Fortaleza", data: "2023-09-10", preço: 450.0},
	}
	ViagemMaisCara(viagens)
}
