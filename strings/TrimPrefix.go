/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimPrefix(s, prefix string) string
//
// Retorna s removendo prefix do início.
// Se s não inicia com prefix não é alterada.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Print(s)
}

// Saída:
// Gophers!!!
//
