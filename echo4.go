// Implementação do comando echo. Versão 4.
// arataca89@gmail.com
// 20210412
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}

/////////////////////////////////////////////////////////////////////
// Aqui fmt.Println() formata os.Args
// A saída será os valores de os.Args envolvidos em colchetes.
//
// (DONOVAN e KERNIGHAN, 2017)
