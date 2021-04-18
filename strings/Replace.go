/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Replace(s, old, new string, n int) string
//
// Retorna uma cópia de s substituindo old por new n vezes.
// Se old é "" new é inserido n vezes, sendo na posição 0 e depois de
// cada caracter.
// Se n < 0 todas as ocorrências de old são substituídas.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "", "-moo-", 5))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}

// Saída:
// oinky oinky oink
// -moo-o-moo-i-moo-n-moo-k-moo- oink oink
// moo moo moo
