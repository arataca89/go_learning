// Implementação do comando echo. Versão 1.
// arataca89@gmail.com
// 20210412

package main

import (
	"fmt"
	"os"
)

func main() {
	var linha, espaco string
	for i := 1; i < len(os.Args); i++ {
		linha += espaco + os.Args[i]
		espaco = " "
	}
	fmt.Println(linha)
}

/////////////////////////////////////////////////////////////////////
// O pacote "os" oferece recursos para lidar com o sistema
// operacional, independente de plataforma.
//
// A variável "Args" faz parte do pacote "os" e armazena os
// argumentos de linha de comando.
//
// "Args" é um "slice".
//
// os.Args[0] armazena o nome do arquivo executado.
//
// os.Args[1] armazena o primeiro argumento de linha de comando;
// os.Args[2] armazena o segundo argumento de linha de comando;
// os.Args[3] armazena o terceiro argumento de linha de comando;
// ..... e assim por diante.
//
// len(os.Args) retorna o número de elementos de os.Args
//
// (DONOVAN e KERNIGHAN, 2017)
