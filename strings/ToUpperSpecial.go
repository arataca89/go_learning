/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ToUpperSpecial(c unicode.SpecialCase, s string) string
//
// Retorna uma cópia de s com todas as letras maiúsculas, conforme
// mapeamento especificado por c.
//

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

// Saída:
// ÖRNEK İŞ
//
