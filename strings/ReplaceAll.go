/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ReplaceAll(s, old, new string) string
//
// Retorna uma cópia de s substituindo todas as ocorrências de old
// por new.
// Se old é "" new é inserido na posição 0 e depois de cada caracter.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	fmt.Println(strings.ReplaceAll("oink oink oink", "", "-moo-"))
}

// Saída:
// moo moo moo
// -moo-o-moo-i-moo-n-moo-k-moo- -moo-o-moo-i-moo-n-moo-k-moo- -moo-o-moo-i-moo-n-moo-k-moo-
