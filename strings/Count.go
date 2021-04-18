/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Count(s, substr string) int
//
// Retorna o número de ocorrências de string em s.
// Se substr for "" retorna 1 + o número de Unicode code points em s.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("cheese", "z")) // 0
	fmt.Println(strings.Count("five", ""))    // 5
}
