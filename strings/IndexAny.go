/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func IndexAny(s, chars string) int
//
// Retorna o índice da primeira ocorrência de qualquer dos caracteres
// existentes em chars.
// se não houver nenhuma ocorrência retorna -1.
//
// Fonte: https://golang.org/pkg/strings/
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1
}
