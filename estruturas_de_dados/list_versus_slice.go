/********************************************************************
list versus slice
arataca89@gmail.com
20210504

Este programa apresenta o tempo gasto por uma list e uma slice
para inserir 2 itens no final e 20 itens no início.

Após execução observa-se que a vesão usando list é aproximadamente
5 vezes mais rápida que a versão usando slice.

Assim, se seu programa precisar inserir ou deletar vários itens em
várias posições o pacote container/list é uma escolha melhor que o
slice.

Referência:
https://www.dotnetperls.com/container-list-go
*/
package main

import (
	"container/list"
	"fmt"
	"strconv"
	"time"
)

func main() {
	t0 := time.Now()

	// Usando list.
	for i := 0; i < 10000; i++ {
		lista := list.New()
		// adicionando 2 itens no final
		lista.PushBack("Golang")
		lista.PushBack("Gopher")
		// adicionando 20 itens no início
		for i := 0; i < 20; i++ {
			// Itoa() converte int para string
			lista.PushFront(strconv.Itoa(i))
		}
	}

	t1 := time.Now()

	// Usando slice.
	for i := 0; i < 10000; i++ {
		lista := []string{}
		// adicioando 2 itens no final.
		lista = append(lista, "Golang")
		lista = append(lista, "Gopher")
		// adicionando 20 itens no início.
		for i := 0; i < 20; i++ {
			// Cria um novo slice e adiciona uma string...
			tempSlice := []string{}
			tempSlice = append(tempSlice, strconv.Itoa(i))
			// ... agora adiciona as strings já existens após esta
			for x := range lista {
				tempSlice = append(tempSlice, lista[x])
			}
			lista = tempSlice
		}
	}

	t2 := time.Now()

	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}

/********************************************************************
SAÍDA:

19.7843ms
96.1101ms

*/
