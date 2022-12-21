package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			fmt.Println("exibirLogs")
			imprimeLogs()
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
	fmt.Println("2- exibirLogs")
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

	// sites := []string{"https://www.facebook.com", "https://www.instagram.com", "https://www.americanas.com.br"}

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso.")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "erro.StatusCode:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
