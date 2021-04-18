/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Map(mapping func(rune) rune, s string) string
//
// Retorna uma cópia de s com os caracteres modificados de acordo
// com o mapeamento da função mapping().
//
// Se mapping() retornar um número negativo o caractere será removido
//
// Fonte: https://golang.org/pkg/strings/
//
// Tabela dos caracteres UNICODE:
// https://en.wikipedia.org/wiki/List_of_Unicode_characters
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	vogal := func(r rune) rune {
		if strings.ContainsRune("aeiou", r) {
			return 86
		}
		if r == '.' {
			return -1
		}
		return r
	}
	fmt.Println("Twas brillig and the slithy gopher...")
	fmt.Println(strings.Map(vogal, "Twas brillig and the slithy gopher..."))
}

// Saída:
// Twas brillig and the slithy gopher...
// TwVs brVllVg Vnd thV slVthy gVphVr
