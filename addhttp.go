// Verifica se os argumentos de linha de comando têm o prefixo
// "http://". Se não tiverem, insere.
// arataca89@gmail.com
// 20210413
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, valor := range os.Args[1:] {
		if strings.HasPrefix(valor, "http://") == false {
			fmt.Println("http://" + valor)
		} else {
			fmt.Println(valor)
		}
	}
}
