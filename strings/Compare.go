/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Compare(a, b string) int
//
// Compara duas strings.
// Se a == b retorna 0
// Se a < b retorna -1
// Se a > b retorna +1
//
// OBSERVAÇÃO: É mais rápido usar os operadores internos de
// comparação ==, <, > que usar a função Compare()
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1

	// Usando os operadores internos
	fmt.Println("a" == "a") // true
	fmt.Println("a" == "b") // false
	fmt.Println("a" > "b")  // false
	fmt.Println("a" < "b")  // true
}
