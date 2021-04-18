/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ToValidUTF8(s, replacement string) string
//
// Retorna uma cópia de s com cada sequência de bytes UTF-8 inválidas
// substituídas por replacement.
//

package main

// importing fmt and strings
import (
	"fmt"
	"strings"
)

// calling main method
func main() {

	// '\xc5' é uma sequência UTF-8 inválida
	fmt.Print(strings.ToValidUTF8("Geeks\xc5Geeks", "For"))
}

// Saída:
// GeeksForGeeks
//
