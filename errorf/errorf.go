/********************************************************************
Exemplo de função do tipo errorf(), ou seja, que retorna uma
mensagem de erro.
arataca89@gmail.com
20210505

errorf() monta uma mensagem de erro formatada indicando o número da
linha no início.

O sufixo "f" usado no final do nome da função errorf() é uma
convenção usada para nomes de funções variádicas que aceitam uma
string de formatação no estilo de Printf().

Referências:
(Donovan e Kernighan, 2017)
https://golang.org/pkg/fmt/#Fprintf
*/
package main

import (
	"fmt"
	"os"
)

// Errorf retorna uma mensagem de erro formatada.
func Errorf(lineNum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Linha %d: ", lineNum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	linha := 13
	termo := "atola"
	Errorf(linha, "termo desconhecido: %s", termo)

	linha = 41
	numero := 1185
	Errorf(linha, "número não registrado: %d", numero)

	linha = 111
	termo = "papiro"
	numero = 8
	Errorf(linha, "bizú:%s; por %d horas, no mínimo", termo, numero)
}

/********************************************************************
SAÍDA:

Linha 13: termo desconhecido: atola
Linha 41: número não registrado: 1185
Linha 111: bizú:papiro; por 8 horas, no mínimo

*/
