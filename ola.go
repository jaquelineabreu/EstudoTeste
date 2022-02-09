package main

import "fmt"

func Ola(nome string) string {
	return "Olá, " + nome
}

//regra de negócio
func main() {
	//fmt.Println("Olá, mundo")
	fmt.Println(Ola("mundo"))
}
