/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func ToTitle(s string) string
//
// Retorna uma cópia de s com todas as letras maiúsculas.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare este exemplo com o exemplo da função Title().
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
}

// Saída:
// HER ROYAL HIGHNESS
// LOUD NOISES
// ХЛЕБ
//
