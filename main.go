package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	displaysIntroduction()

	for {

		displaysMenu()

		aplicacao := leAplicacao()

		switch aplicacao {
		case 1:
			iniciarMonitorando()
		case 2:
			fmt.Println("menu")

		case 0:
			fmt.Println("close")
			os.Exit(0)
		default:
			fmt.Println("Aplicação não identificada")
			os.Exit(-1)
		}

	}

}

func displaysIntroduction() {
	var email string = "joaovitorfoca2001@gmail.com"
	var name string = "João Vitor"
	var years string = "21"
	var version float32 = 1.0
	fmt.Println("Email:", email, "\nHello Sr(a)", name, "\nyears", years)
	fmt.Println("version", version)

}

func displaysMenu() {
	fmt.Println("1- monitorando")
	fmt.Println("2- menu")
	fmt.Println("0- close")
}

func leAplicacao() int {
	var aplicacaoLida int
	fmt.Scan(&aplicacaoLida)
	fmt.Println("A Aplicação selecionada", aplicacaoLida)
	fmt.Println("")

	return aplicacaoLida
}

func iniciarMonitorando() {
	fmt.Println("monitorando")
	sites := []string{"https://www.facebook.com", "https://www.instagram.com", "https://www.americanas.com.br"}
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("testando site:", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso.")
	} else {
		fmt.Println("Site:", site, "erro.StatusCode:", resp.StatusCode)
	}
}
