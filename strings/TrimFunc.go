/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimFunc(s string, f func(rune) bool) string
//
// Retorna s removendo os caracteres especificados conforme a função
// f() do início e do final.
//

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}

// Saída:
// Hello, Gophers
//
