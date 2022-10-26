package main

import "fmt"

func main() {
	oi := seunome("Hiago")
	fmt.Println(oi)
}

func seunome(nome string) (resposta string) {
	resposta = fmt.Sprintf("Bem vindo " + nome)
	return
}
