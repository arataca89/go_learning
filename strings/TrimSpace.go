/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimSpace(s string) string
//
// Retorna s removendo os caracteres de espaço do início e do final.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}

// Saída:
// Hello, Gophers
//
