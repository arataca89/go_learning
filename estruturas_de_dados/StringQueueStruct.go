/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210426
//
// Implementação uma fila de strings usando um tipo definido pelo
// usuário e ponteiros.
//
// Referências:
// https://br.ccm.net/faq/10219-as-filas-em-linguagem-c

package main

import (
	"fmt"
	"os"
)

type Node struct {
	data string
	next *Node
	prev *Node
}

var start *Node
var end *Node
var size int

func printQueue() {
	current := end
	for i := 1; i <= size; i++ {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func push(item string) {
	newNode := new(Node)
	if size == 0 {
		newNode.data = item
		newNode.next = nil
		newNode.prev = nil
		start = newNode
		end = newNode
		size++
	} else {
		newNode.data = item
		newNode.next = end
		end.prev = newNode
		end = newNode
		size++
	}
}

func pop() *Node {
	if size == 0 {
		return nil
	}
	temp := start
	start = start.prev
	size--
	return temp
}

func main() {
	var i string

	for {
		fmt.Println("<< 1 >> Inserir item na fila")
		fmt.Println("<< 2 >> Retirar item da fila")
		fmt.Println("<< 3 >> Exibir fila")
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
				fmt.Printf("\nFila vazia\n\n")
			} else {
				fmt.Println("\nItem retirado:", itemPop.data)
				fmt.Println()
			}
		} else if i == "3" {
			fmt.Printf("\nExibindo fila. Quantidade de itens: %d\n", size)
			fmt.Println("---------------------------------------")
			printQueue()
			fmt.Printf("\n")
		} else {
			fmt.Printf("\nOpção inválida!\n\n")
		}
	}
}
