package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresFavoritos []string
}

func main() {

	pessoa1 := Pessoa{
		nome:             "Matheus",
		sobrenome:        "Santos",
		saboresFavoritos: []string{"morango", "chocolate"},
	}
	fmt.Println(pessoa1)
}
