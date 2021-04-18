/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func TrimSuffix(s, suffix string) string
//
// Retorna s removendo suffix do final.
// Se s não termina com suffix não é alterada.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Print(s)
}

// Saída:
// ¡¡¡Hello
//
