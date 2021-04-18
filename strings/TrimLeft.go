/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimLeft(s, cutset string) string
//
// Retorna s removendo os caracteres especificados em cutset
// do início.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}

// Saída:
// Hello, Gophers!!!
//
