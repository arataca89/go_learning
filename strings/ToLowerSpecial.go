/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ToLowerSpecial(c unicode.SpecialCase, s string) string
//
// Retorna s com todas as letras minúsculas conforme o mapeamento
// especificado por c.
//

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}

// Saída:
// önnek iş
//
