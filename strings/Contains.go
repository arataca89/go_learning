/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Contains(s, substr string) bool
//
// Se s contém substr retorna true, senão retorna false
//
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
}
