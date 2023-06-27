package main

import "fmt"

type musicas struct {
	Titulo  string
	Artista string
	Duração int
}

type Playlist struct {
	Nome    string
	Musicas []musicas
}

func sons(musica Playlist) {
	fmt.Println("O nome da playlista é:", musica.Nome)
	tempo_playlist := 0

	for _, sons := range musica.Musicas {
		fmt.Printf("Musica: %s\nArtista:%s\nDuração:%d minutos\n", sons.Titulo, sons.Artista, sons.Duração)
		tempo_playlist += sons.Duração
	}
	fmt.Printf("E o tempo total da playlist é de %d minutos", tempo_playlist)
}

func main() {
	musicas := []musicas{
		{Titulo: "Nao faz isso comigo nao", Artista: "Tz", Duração: 3},
		{Titulo: "Materia", Artista: "Leviano", Duração: 4},
		{Titulo: "Mítico Joven", Artista: "Ryu, CORREDOR", Duração: 3},
	}
	playlist1 := Playlist{
		Nome:    "Trap",
		Musicas: musicas,
	}

	sons(playlist1)
}
