/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func LastIndexByte(s string, c byte) int
//
// Retorna o índice da última ocorrência de c em s ou -1 se não
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
	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) // 10
	fmt.Println(strings.LastIndexByte("Hello, world", 'o')) // 8
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) // -1
}
