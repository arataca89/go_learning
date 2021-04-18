/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func HasPrefix(s, prefix string) bool
//
// Retorna true se a string s inicia com prefix.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "C"))  // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true
}
