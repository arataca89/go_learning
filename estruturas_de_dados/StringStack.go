/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210418
//
// Implementação uma pilha de strings
//
// func new(Type) *Type
//
// new() aloca memória. Recebe um tipo e retorna um ponteiro para
// o tipo passado como argumento.
//
// Referência:
// (YANG HU, 2020)
//

package main

import (
	"fmt"
	"os"
)

type itemStack struct {
	data string     // dados; string a ser armazenada na pilha
	next *itemStack // ponteiro para próximo item na pilha
}

var top *itemStack = nil // aponta para topo da pilha
var size int             // tamanho da pilha

func push(item string) {
	if top == nil {
		top = new(itemStack)
		top.data = item
	} else {
		var newItem *itemStack = new(itemStack)
		newItem.data = item
		newItem.next = top
		top = newItem
	}
	size++
}

func pop() *itemStack {
	if top == nil {
		return nil
	}

	var ptr = top
	top = top.next

	ptr.next = nil
	size--

	return ptr
}

func printStack() {
	var itemCurrent = top
	for i := size; i > 0; i-- {
		fmt.Println(itemCurrent.data)
		itemCurrent = itemCurrent.next
	}
}

func main() {
	var i string

	for {
		fmt.Println("<< 1 >> Inserir item na pilha")
		fmt.Println("<< 2 >> Retirar item da pilha")
		fmt.Println("<< 3 >> Exibir pilha")
		fmt.Println("<< 0 >> Sair")

		fmt.Print("Entre com sua opção: ")
		fmt.Scanf("%s\r", &i)

		if i == "0" {
			os.Exit(0)
		} else if i == "1" {
			fmt.Print("Entre com a string a ser inserida: ")
			fmt.Scanf("%s\r", &i)
			push(i)
			fmt.Println()
		} else if i == "2" {
			itemPop := pop()
			if itemPop == nil {
				fmt.Printf("\nPilha vazia\n\n")
			} else {
				fmt.Println("\nItem retirado", itemPop.data)
				fmt.Println()
			}
		} else if i == "3" {
			fmt.Println("\nExibindo a pilha")
			fmt.Println("----------------")
			printStack()
			fmt.Println()
		} else {
			fmt.Printf("\nOpção inválida!\n\n")
		}
	}
}

/////////////////////////////////////////////////////////////////////
/*
TESTE

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 1
Entre com a string a ser inserida: um

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 1
Entre com a string a ser inserida: dois

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 1
Entre com a string a ser inserida: tres

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 3

Exibindo a pilha
----------------
tres
dois
um

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 2

Item retirado tres

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 3

Exibindo a pilha
----------------
dois
um

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 2

Item retirado dois

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 3

Exibindo a pilha
----------------
um

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 2

Item retirado um

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 3

Exibindo a pilha
----------------

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 2

Pilha vazia

<< 1 >> Inserir item na pilha
<< 2 >> Retirar item da pilha
<< 3 >> Exibir pilha
<< 0 >> Sair
Entre com sua opção: 0

*/
