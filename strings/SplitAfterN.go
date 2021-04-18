/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func SplitAfterN(s, sep string, n int) []string
//
// Retorna um slice dos tokens de s separados por sep mas incluem
// sep.
// n determina a quantidade de tokens.
// n > 0 	no máximo n tokens
// n == 0 	0 tokens
// n < 0	todos os tokens
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.SplitAfterN("a,b,c,d,e", ",", 3))
	fmt.Println(strings.SplitAfterN("a,b,c,d,e", ",", 0))
	fmt.Println(strings.SplitAfterN("a,b,c,d,e", ",", -1))
}

// Saída:
// [a, b, c,d,e]
// []
// [a, b, c, d, e]
