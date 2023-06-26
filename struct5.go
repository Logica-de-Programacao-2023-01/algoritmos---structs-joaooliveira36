package main

import (
	"bufio"
	"fmt"
	"os"
)

type Musicas struct {
	Titulo  string
	Artista string
	Duração int
}

type Minha_Playlist struct {
	Nome    string
	Musicas []Musicas
}

func Sons(musica Minha_Playlist) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual música você deseja?")
	scanner.Scan()
	musica_desejada := scanner.Text()

	for _, sons := range musica.Musicas {
		if sons.Titulo == musica_desejada {
			fmt.Printf("Musica: %s\nArtista:%s\nDuração:%d minutos\n", sons.Titulo, sons.Artista, sons.Duração)
		}
	}
}

func main() {
	musicas := []Musicas{
		{Titulo: "Nao faz isso comigo nao", Artista: "Tz", Duração: 3},
		{Titulo: "Matéria", Artista: "Leviano", Duração: 4},
		{Titulo: "Mítico Jovem", Artista: "Ryu, CORREDOR", Duração: 3},
	}
	playlist1 := Minha_Playlist{
		Nome:    "Trap",
		Musicas: musicas,
	}

	Sons(playlist1)
}
