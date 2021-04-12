// Implementação do comando echo. Versão 3.
// arataca89@gmail.com
// 20210412
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/////////////////////////////////////////////////////////////////////
// Nesta versão foi usada a função strings.Join()
//
// (DONOVAN e KERNIGHAN, 2017)
