/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Title(s string) string
//
// Retorna uma cópia de s com todas as primeiras letras maiúsculas.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare este exemplo com o exemplo da função ToTitle().
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))
	fmt.Println(strings.Title("хлеб"))
}

// Saída:
// Her Royal Highness
// Loud Noises
// Хлеб
//
