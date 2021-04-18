/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Split(s, sep string) []string
//
// Retorna um slice dos tokens de s separados por sep.
// Se s não contém sep e sep não é vazio, retorna um slice de tamanho
// 1 cujo único elemento é s.
// Se sep é "" retorna um slice separando cada caracter de s.
// Se s e sep são iguais a "" retorna uma slice vazia.
// É equivalente a SplitN com n < 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("Dividindo em tokens", " "))
	fmt.Println(strings.Split("a,e,i,o,u", ","))
	fmt.Println(strings.Split("I;II;III;IV;V", ";"))
	fmt.Println(strings.Split("Dividindo em tokens", "#"))
	fmt.Println(strings.Split("a,e,i,o,u", "#"))
	fmt.Println(strings.Split("Dividindo em tokens", ""))
	fmt.Println(strings.Split("", ""))

}

// Saída:
// [Dividindo em tokens]
// [a e i o u]
// [I II III IV V]
// [Dividindo em tokens]
// [a,e,i,o,u]
// [D i v i d i n d o   e m   t o k e n s]
// []
