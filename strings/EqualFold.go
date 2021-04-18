/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func EqualFold(s, t string) bool
//
// Retorna true se s e t forem iguais independente se as letras
// forem maiúsculas ou minúsculas
//
// Fonte: https://golang.org/pkg/strings/
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Go", "go"))         // true
	fmt.Println(strings.EqualFold("GoLang", "golANG")) // true
}
