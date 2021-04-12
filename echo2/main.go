// Implementação do comando echo. Versão 2.
// arataca89@gmail.com
// 20210412

package main

import (
	"fmt"
	"os"
)

func main() {
	linha, espaco := "", ""

	for _, token := range os.Args[1:] {
		linha += espaco + token
		espaco = " "
	}

	fmt.Println(linha)
}

/////////////////////////////////////////////////////////////////////
// _, token := range os.Args[1:]
//
// "range" no comando acima retorna dois valores: o índice e o valor
// de os.Args
//
// _ (underline ou underscore, sublinhado) é chamado
// "identificador vazio" e deve ser usado sempre que uma variável
// for exigida na sintaxe da linguagem Go mas o programador não
//  quiser declarar uma variável só para essa finalidade.
//
// := (dois pontos seguido de igualdade) define  a forma curta de
// declaração de variáveis. Ela declara e inicializa a variável
// ao mesmo tempo. Só pode ser usada dentro de funções.
//
// (DONOVAN e KERNIGHAN, 2017)
