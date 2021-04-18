/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func LastIndexAny(s, chars string) int
//
// Retorna o índice da última ocorrência de qualquer caracter de
// chars existente em s ou -1 se não houver ocorrência.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexAny("go gopher", "go"))     // 4
	fmt.Println(strings.LastIndexAny("go gopher", "rodent")) // 8
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))   // -1
}
