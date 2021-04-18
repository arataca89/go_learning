/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func FieldsFunc(s string, f func(rune) bool) []string
//
// Separa a string s em tokens e retorna um slice com estes tokens.
// Os separadores dos tokens são configurados conforme a função f
//
// No exemplo dado observe que, no retorno de f, se o caracter
// NÃO for uma letra E NÃO for um número a função retorna true
// indicando que o caracter é um separador.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("  foo1;bar2,baz3...", f)) // [foo1 bar2 baz3]

	fmt.Println(strings.FieldsFunc("  ", f)) // []
}
