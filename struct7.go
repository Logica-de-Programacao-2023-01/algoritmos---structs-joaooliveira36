package main

import "fmt"

type Animal struct {
	nome    string
	espécie string
	idade   string
	som     string
}

func sons_animal(animal Animal) {
	var som_desejado string
	fmt.Print("Qual som voce deseja que o animal faça?")
	fmt.Scan(&som_desejado)

	animal.som = som_desejado

	fmt.Print("O nome do animal é " + animal.nome + " é da especie " + animal.espécie + " tem " + animal.idade +
		" anos " + "e faz o barulho " + animal.som)
}

func main() {
	animal1 := Animal{
		nome:    "Bruno",
		espécie: "cachorro",
		idade:   "6",
		som:     "AuAu",
	}
	sons_animal(animal1)
}
