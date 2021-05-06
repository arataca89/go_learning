/*
O pacote sort1 demonstra o uso básico do pacote sort para ordenar
estruturas de dados definidas pelo usuário.

O pacote sort possui funções que facilitam a ordenação de slices e
tipos definidos pelo usuário.

Referências:
Doxsey(2016)
https://golang.org/pkg/sort/
https://golang.org/pkg/sort/#Interface

arataca89@gmail.com
20210506
*/
package main

import (
	"fmt"
	"sort"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) String() string {
	return fmt.Sprintf("%s: %d", p.Nome, p.Idade)
}

// PorIdade implementa sort.Interface para []Pessoa baseado
// no campo Idade.
type PorIdade []Pessoa

// A interface sort.Interface exige a definição das funções abaixo
// conforme: https://golang.org/pkg/sort/#Interface
func (a PorIdade) Len() int           { return len(a) }
func (a PorIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorIdade) Less(i, j int) bool { return a[i].Idade < a[j].Idade }

func main() {
	personagem := []Pessoa{
		{"Woody", 13},
		{"Buzz", 41},
		{"Rex", 11},
		{"Zurg", 85},
	}

	fmt.Println("slice desordenado: ")
	fmt.Println(personagem)
	fmt.Println()

	// Existem duas maneiras de ordenar uma slice. Primeiro, pode-se
	// definir um conjunto de métodos para o tipo slice, como
	// PorIdade, e chamar sort.Sort. Este primeiro exemplo usa esta
	// técnica.
	fmt.Println("Ordenando pela idade usando sort.Sort(): ")
	sort.Sort(PorIdade(personagem))
	fmt.Println(personagem)
	fmt.Println()

	// Outra maneira é usar sort.Slice com uma função customizada
	// Less, a qual pode ser implementada como uma closure.
	// Neste caso nenhum método é necessário. Se os métodos existirem
	// serão ignorados. Aqui a slice foi reordenada na ordem reversa:
	// compare a função closure com PorIdade.Less.
	fmt.Println("Ordenando pela idade, na ordem inversa, usando sort.Slice(): ")
	sort.Slice(personagem, func(i, j int) bool {
		return personagem[i].Idade > personagem[j].Idade
	})
	fmt.Println(personagem)

}

/********************************************************************

SAÍDA:

slice desordenado:
[Woody: 13 Buzz: 41 Rex: 11 Zurg: 85]

Ordenando pela idade usando sort.Sort():
[Rex: 11 Woody: 13 Buzz: 41 Zurg: 85]

Ordenando pela idade, na ordem inversa, usando sort.Slice():
[Zurg: 85 Buzz: 41 Woody: 13 Rex: 11]

*/
