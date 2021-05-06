/*
O pacote ponteiros1 demonstra o uso básico de ponteiros em Go.

Um ponteiro é um endereço de memória.

O operador * declara um tipo ponteiro:
	var nrPtr *int

O operador & retorna o endereço de uma variável:
	var nr int
	&nr retorna o endereço de memória da variável nr.

Para acessar o valor armazenado na memória, apontado por um ponteiro,
Use o operador * antes do ponteiro:
	*nrPtr retorna o valor armazenado em nrPtr (caso nrPtr seja um
	ponteiro)

arataca89@gmail.com
20210506
*/
package main

import "fmt"

func main() {

	var nr int = 13

	// *int declara a variável nrPtr como um ponteiro para um int
	// Um ponteiro armazena o endereço de memória de uma variável
	// O operador & retorna o endereço de memória da variável
	var nrPtr *int = &nr

	fmt.Println("nr:", nr)       // nr: 13
	fmt.Println("nrPtr:", nrPtr) // nrPtr: 0xc0000ac058

	// o operador * antes de um ponteiro acessa o valor armazenado
	// no endereço apontado.
	fmt.Println("*nrPtr: ", *nrPtr) // *nrPtr:  13
}
