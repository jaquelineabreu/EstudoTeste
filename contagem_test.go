package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Go!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Pausa() {
	s.Chamadas++
}
