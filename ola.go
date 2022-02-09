package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoOlaPortugues + nome
}

//regra de negócio
func main() {
	//fmt.Println("Olá, mundo")
	fmt.Println(Ola("mundo"))
}
