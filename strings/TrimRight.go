/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimRight(s, cutset string) string
//
// Retorna s removendo os caracteres especificados em cutset
// do final.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
}

// Saída:
// ¡¡¡Hello, Gophers
//
