/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func IndexRune(s string, r rune) int
//
// Retorna o índice da primeira ocorrência de r em s ou -1 se
// não houver ocorrência.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1
}
