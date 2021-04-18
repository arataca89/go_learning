/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Index(s, substr string) int
//
// Retorna o índice da primeira ocorrência de substr em s
// ou -1 se s não possuir substr.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1
}
