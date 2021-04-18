/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimRightFunc(s string, f func(rune) bool) string
//
// Retorna s removendo os caracteres especificados conforme a função
// f() do final.
//

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}

// Saída:
// ¡¡¡Hello, Gophers
//
