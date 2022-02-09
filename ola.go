package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixodeSaudacao(idioma) + nome

}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

// func Ola(nome string, idioma string) string {
// 	if nome == "" {
// 		nome = "mundo"
// 	}

// 	if idioma == espanhol {
// 		return prefixoOlaEspanhol + nome
// 	}

// 	if idioma == frances {
// 		return prefixoOlaFrances + nome
// 	}

// 	return prefixoOlaPortugues + nome
// }

//regra de negócio
func main() {
	//fmt.Println("Olá, mundo")
	fmt.Println(Ola("Jaque", frances))
}
