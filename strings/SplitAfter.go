/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func SplitAfter(s, sep string) []string
//
// Retorna um slice dos tokens de s separados por sep mas incluem
// sep.
// Se s não contém sep e sep não é vazio, retorna um slice de tamanho
// 1 cujo único elemento é s, incluindo sep.
// Se sep é "" retorna um slice separando cada caracter de s, inclui sep.
// Se s e sep são iguais a "" retorna uma slice vazia.
// É equivalente a SplitAfterN com n < 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", "#"))
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ""))
	fmt.Printf("%q\n", strings.SplitAfter("", ""))
}

// Saída:
// ["a," "b," "c"]
// ["a,b,c"]
// ["a" "," "b" "," "c"]
// []
