/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Fields(s string) []string
//
// Separa a string s em tokens e retorna um slice com estes tokens.
// Considera que o separador dos tokens é o caracter de espaço
// definido por unicode.IsSpace
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Fields("  foo bar  baz   ")) // [foo bar baz]
	fmt.Println(strings.Fields("  "))                // []
}
