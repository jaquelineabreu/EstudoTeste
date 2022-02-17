package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper te permite definir pausas
type Sleeper interface {
	Pausa()
}

// SleeperPadrao é uma implementação de Sleeper com um atraso pré-definido
type SleeperPadrao struct{}

// Pausa vai pausar a execução pela Duração definida
func (d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}

const ultimaPalavra = "Go!"
const inicioContagem = 3

// Contagem imprime uma contagem de 3 para a saída com um atraso determinado por um Sleeper
func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

const escrita = "escrita"
const pausa = "pausa"
