package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Go!"
const inicioContagem = 3

type SleeperSpy struct {
	Chamadas int
}

type Sleeper interface {
	Sleep()
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
