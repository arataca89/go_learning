/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func LastIndex(s, substr string) int
//
// Retorna o índice da última ocorrência de substr em s ou -1 se não
// houver ocorrência.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
}
