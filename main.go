package testego

import "fmt"

//seunome usado para teste.
func Seunome(nome string) (resposta string) {
	resposta = fmt.Sprintf("Bem vindo " + nome)
	return
}
