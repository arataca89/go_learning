/*
O pacote ponteiros2 mostra a diferença entre a passagem por valor
e a passagem por referência

Na passagem por valor uma cópia do dado é passada para a função
e o valor original não é alterado.const

Na passagem por referência é passado o endereço de memória onde o
dado está armazenado. Neste caso as alterações feitas pela função
podem alterar o dado.

arataca89@gmail.com
20210506
*/
package main

import "fmt"

// passagem por valor (dados não são alterados)
func quadrado1(n int) int {
	return n * n
}

// passagem por referência (dados podem ser alterados)
func quadrado2(n *int) {
	*n = *n * *n
}

func main() {

	numero := 3

	fmt.Println("numero antes chamar quadrado1():", numero)
	fmt.Println("Chamando quadrado1(numero):", quadrado1(numero))
	fmt.Println("numero após a chamada de quadrado1():", numero)
	fmt.Println()
	fmt.Println("chamando quadrado2(&numero)")
	quadrado2(&numero)
	fmt.Println("numero após a chamada de quadrado2(&numero)", numero)

}
