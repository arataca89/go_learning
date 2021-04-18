/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210418
//
// Algumas vezes há a necessidade de juntar várias strings para
// formar uma string maior. Este operação pode provocar problemas
// de performance.
// Neste caso o tipo Builder pode ser usado para isso.
// Com ele podemos juntar strings de maneira rápida e com boa
// performance. Poucas alocações são necessárias.
//
// Referências:
// https://www.dotnetperls.com/strings-builder-go
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Cria novo Builder.
	builder := strings.Builder{}

	// Escreve uma string 5 vezes.
	// WriteString() adiciona a string ao Builder
	for i := 0; i < 5; i++ {
		builder.WriteString("Builder ")
	}

	// Converte Builder em String e exibe no console.
	// String() converte os dados buferizados no Builder numa string
	result := builder.String()
	fmt.Println(result)
}

// Saída:
// Builder Builder Builder Builder Builder
//
