/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Join(elems []string, sep string) string
//
// Junta os itens de elems numa única string separados por sep.
//
// Fonte: https://golang.org/pkg/strings/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, " ### "))
}

// Saída: foo ### bar ### baz
