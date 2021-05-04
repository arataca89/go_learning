/********************************************************************
container/list uso básico
arataca89@gmail.com
20210504

O pacote container/list implementa uma lista duplamente ligada.

Cada item da lista possui um ponteiro para o próximo item e um
ponteiro para o item anterior. Isto torna a inserção e remoção
de itens mais rápida mas retarda a iteração.

Permite a construção de programas que usam listas com uma
melhor performance.

Referências:
Caleb Doxsey, 2016, Introdução a linguagem Go
https://golang.org/pkg/container/list/
https://www.dotnetperls.com/container-list-go
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	// criando uma lista
	lista := list.New()

	// PushBack() insere um item no final da lista
	// É similar ao método append() usado com slices.
	lista.PushBack("Buzz")
	lista.PushBack("Woody")
	lista.PushFront("Zurg")

	fmt.Println("Lista com os três primeiros itens:")

	// Iterando através da lista.
	// Os métodos Front() e Next() são usados para iterar através
	// da lista. A iteração através de uma lista é mais lenta que
	// através de um slice ou de um array.
	for item := lista.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
	fmt.Println()

	fmt.Println("Adicionando itens no início da lista:")

	// PushFront() insere um item no início da lista.
	// Este método mostra uma vantagem do uso das listas
	// pois inserir um item no início é rápido.
	lista.PushFront("Jessie")
	lista.PushFront("Rex")
	lista.PushFront("Slinky")

	for item := lista.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
	fmt.Println()

}

/********************************************************************

SAÍDA:

Lista com os três primeiros itens:
Zurg
Buzz
Woody

Adicionando itens no início da lista:
Slinky
Rex
Jessie
Zurg
Buzz
Woody

*/
