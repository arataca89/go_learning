// Implementação do comando echo. Versão 5.
// arataca89@gmail.com
// 20210412

package main

import (
	"fmt"
	"os"
)

func main() {
	linha, espaco := "", ""

	fmt.Println("File name: ", os.Args[0])

	fmt.Print("arguments: ")
	for _, token := range os.Args[1:] {
		linha += espaco + token
		espaco = " "
	}

	fmt.Println(linha)
}

/////////////////////////////////////////////////////////////////////
// os.Args[0] exibe o nome do arquivo.
