/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func IndexByte(s string, c byte) int
//
// Retorna o índice da primeira ocorrência de c em s ou -1 se não
// houver ocorrência.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexByte("golang", 'g'))  // 0
	fmt.Println(strings.IndexByte("gophers", 'h')) // 3
	fmt.Println(strings.IndexByte("golang", 'x'))  // -1
}
