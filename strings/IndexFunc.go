/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func IndexFunc(s string, f func(rune) bool) int
//
// Retorna o índice da primeira ocorrência de caracter que satisfaz a
// função f() ou -1 se não houver ocorrência.
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
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}
