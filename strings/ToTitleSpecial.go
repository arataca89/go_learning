/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ToTitleSpecial(c unicode.SpecialCase, s string) string
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
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}

// Saída:
// DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
//
