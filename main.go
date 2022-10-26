package testego

import "fmt"

//seunome usado para teste.
func seunome(nome string) (resposta string) {
	resposta = fmt.Sprintf("Bem vindo " + nome)
	return
}
