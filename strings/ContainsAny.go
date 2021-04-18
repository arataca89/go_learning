/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ContainsAny(s, chars string) bool
//
// Se s contém algum caractere da string chars retorna true,
// senão retorna false
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Team", "T"))   // true
	fmt.Println(strings.ContainsAny("Team", "e"))   // true
	fmt.Println(strings.ContainsAny("Team", "a"))   // true
	fmt.Println(strings.ContainsAny("Team", "m"))   // true
	fmt.Println(strings.ContainsAny("Team", "z"))   // false
	fmt.Println(strings.ContainsAny("Team", "eam")) // true
	fmt.Println(strings.ContainsAny("Team", "exz")) // true
	fmt.Println(strings.ContainsAny("Team", "xyz")) // false
}
