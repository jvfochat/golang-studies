package main

import (
	"fmt"
	"os"
)

func main() {

	displaysIntroduction()
	displaysMenu()

	aplicacao := leAplicacao()

	switch aplicacao {
	case 0:
		fmt.Println("home")
	case 1:
		fmt.Println("menu")
	case 2:
		fmt.Println("close")
		os.Exit(2)
	default:
		fmt.Println("Aplicação não identificada")
		os.Exit(-1)
	}
}

func displaysIntroduction() {
	var email string = "joaovitorfoca2001@gmail.com"
	var name string = "João Vitor"
	var birthdate string = "18/05/2001"
	var passoword int = 1234
	var version float32 = 1.0
	fmt.Println(email, "\nHello Sr(a).", name, "\nBirth Date:", birthdate, "\nPassoword:", passoword)
	fmt.Println("version", version)

}

func displaysMenu() {
	fmt.Println("0- home")
	fmt.Println("1- menu")
	fmt.Println("2- close")
}

func leAplicacao() int {
	var aplicacaoLida int
	fmt.Scanf("%d", &aplicacaoLida)
	fmt.Println("A Aplicação selecionada", aplicacaoLida)

	return aplicacaoLida
}
