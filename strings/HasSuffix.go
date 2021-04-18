/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func HasSuffix(s, suffix string) bool
//
// Retorna true se a string s termina com suffix.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
	fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
	fmt.Println(strings.HasSuffix("Amigo", ""))    // true
}
