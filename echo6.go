// Implementação do comando echo. Versão 6.
// arataca89@gmail.com
// 20210412

package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, ": ", os.Args[i])
	}
}

/////////////////////////////////////////////////////////////////////
// Esta versão exibe o índice e o valor de cada item em
// os.Args, um por linha.
//
// Exemplo:
//
// go run main.go testando 1 dois III
//
// 0 :  C:\Users\nerd\b001\exe\main.exe
// 1 :  testando
// 2 :  1
// 3 :  dois
// 4 :  III
