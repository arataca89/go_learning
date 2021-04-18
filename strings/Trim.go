/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Trim(s, cutset string) string
//
// Retorna s removendo os caracteres especificados em cutset
// do início e do final.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}

// Saída:
// Hello, Gophers
//
