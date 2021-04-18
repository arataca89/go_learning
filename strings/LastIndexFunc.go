/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func LastIndexFunc(s string, f func(rune) bool) int
//
// Retorna o índice da última ocorrência do caracter que satisfaz
// a função f() ou -1 se não houver ocorrência.
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
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) // 5
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) // 2
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))     // -1
}
