/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210420
//
// Implementação uma pilha de strings usando o tipo slice
//
// Referência:
// (DONOVAM e KERNIGHAN, 2017)
//

package main

import (
	"fmt"
	"os"
)

var stack = make([]string, 1)

func push(item string) {
	stack = append(stack, item)
}

func pop() {
	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}
}

func printStack() {
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
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
			if len(stack) == 0 {
				fmt.Printf("\nPilha vazia\n\n")
			} else {
				top := stack[len(stack)-1]
				pop()
				fmt.Println("\nItem retirado", top)
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
