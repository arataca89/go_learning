// exibe os argumentos da linha de comando
// arataca89@gmail.com
// 20210413
package main

import (
	"fmt"
	"os"
)

func main() {
	for indice, valor := range os.Args[:] {
		fmt.Println(indice, " - ", valor)
	}
}

/////////////////////////////////////////////////////////////////////
//
// Exemplo do uso:
//
// go run args1.go one II three IV cinco
// 0  -  C:\Users\nerd\golearning\args1.exe
// 1  -  one
// 2  -  II
// 3  -  three
// 4  -  IV
// 5  -  cinco
//
