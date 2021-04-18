/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ContainsRune(s string, r rune) bool
//
// Se s contém r retorna true, senão retorna false
// r é um Unicode code point
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	// O code point da letra "a", por exemplo, é 97.
	fmt.Println(strings.ContainsRune("aardvark", 97)) // true
	fmt.Println(strings.ContainsRune("timeout", 97))  // false
}
