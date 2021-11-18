// testes de unidade
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoRecebido string
	enderecoEsperado string
}

func TestTipoDeEnderecos(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia ETC", "Rodovia"},
		{"Praça das rosas", "Tipo invalido"},
		{"Estrada teste", "Estrada"},
		{"", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoRecebido := TipoDeEnderecos(cenario.enderecoRecebido)
		if tipoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.enderecoEsperado, tipoRecebido)
		}
	}

}
