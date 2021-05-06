/*
O pacote struct1 demonstra como criar estruturas de dados definidas
pelo usuário e como declarar variáveis do tipo da nova
estrutura criada.

Uma estrutura (struct) é um tipo de dados definido pelo usuário
e é composto por um ou mais campos de dados relacionados.

Pode ser comparada a uma classe sem o recurso da herança, mas com
o recurso da composição.

Referências:
Donovan e Kernighan(2017)
Doxsey(2016)
https://zetcode.com/golang/struct/
https://www.geeksforgeeks.org/structures-in-golang/

arataca89@gmail.com
20210506
*/
package main

import "fmt"

// Define o tipo livro como uma estrutura(struct).
type livro struct {
	titulo    string
	nrPaginas int
	preco     float64
}

func main() {

	// Declarando uma variável do tipo livro.
	var l1 livro
	// Como nada foi atribuído, l1 tem o valor zero em todos os campos
	fmt.Println(l1) // { 0 0}

	// Outra maneira de declarar uma variável struct.
	// Usando esta forma deve-se declarar todos os campos e na ordem
	// em que são declarados na struct.
	var l2 = livro{"A Linguagem de Programação Go", 478, 66.30}
	fmt.Println(l2) // {A Linguagem de Programação Go 478 66.3}

	// Outra maneira de declarar.
	// Aqui também deve-se declarar todos os campos e na ordem correta.
	l3 := livro{"Introdução a linguagem Go", 136, 32.07}
	fmt.Println(l3) // {Introdução a linguagem Go 136 32.07}

	// Outra maneira de declarar.
	// Aqui temos cada campo nomeado seguido de dois pontos e
	// do valor (campo: valor).
	// Oberve a vírgula após o último valor de campo, antes
	// da chave que encerra a declaração da struct. Esta
	// vígula é obrigatória neste tipo de declaração.
	l4 := livro{
		titulo:    "Estrutura de Dados Usando C",
		nrPaginas: 884,
		preco:     292.09,
	}
	fmt.Println(l4) // {Estrutura de Dados Usando C 884 292.09}

	// Neste tipo de declaração não há necessidade de declarar todos
	// os campos. Campos não declarados assumirão o valor zero para
	// o tipo.
	l5 := livro{
		titulo: "Algoritmos em Linguagem C",
		preco:  156.99,
	}
	fmt.Println(l5) // {Algoritmos em Linguagem C 0 156.99}

}
