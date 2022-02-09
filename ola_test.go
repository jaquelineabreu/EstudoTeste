package main

import "testing"

func TestOla(t *testing.T) {
	// resultado := Ola()
	// esperado := "Olá, mundo"

	resultado := Ola("Chris")
	esperado := "Olá, Chris"

	//Teste que que não deve passar

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

}
