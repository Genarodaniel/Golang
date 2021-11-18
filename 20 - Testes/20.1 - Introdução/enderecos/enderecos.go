package enderecos

import "strings"

//TipoDeEnderecos verifica se o endereco tem um tipo válido e o retorna
func TipoDeEnderecos(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoComMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoComMinuscula, " ")[0]

	var enderecoEvalido bool

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoEvalido = true
		}
	}

	if enderecoEvalido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo invalido"
}
