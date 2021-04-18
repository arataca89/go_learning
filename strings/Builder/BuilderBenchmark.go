/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210418
//
// Comparando a concatenação de strings usando o tipo Builder e
// usando string com o operador +.
//
// Referências:
// https://www.dotnetperls.com/strings-builder-go
//

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()

	// Versão 1: usa Builder e WriteString.
	for i := 0; i < 100000; i++ {
		builder := strings.Builder{}
		for v := 0; v < 100; v++ {
			builder.WriteString("Builder")
		}
		if builder.Len() == 0 {
			break
		}
	}

	t1 := time.Now()

	// Versão 2: usa string e o operador +.
	for i := 0; i < 100000; i++ {
		result := ""
		for v := 0; v < 100; v++ {
			result += "Builder"
		}
		if len(result) == 0 {
			break
		}
	}

	t2 := time.Now()
	// Resultados.
	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}

// Saída:
// 83.3855ms
// 1.195975s
//
