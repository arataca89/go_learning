/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210423
//
// Implementa uma fila de strings usando slice
//

package main

import (
	"fmt"
	"os"
)

var queue = make([]string, 1)
var start = true

func push(item string) {
	if start {
		queue[0] = item
		start = false
	} else {
		queue = append(queue, item)
	}
}

func pop() {
	if start {
		fmt.Println("Fila vazia!")
	} else {
		if len(queue) > 0 {
			queue = queue[1:]
		}
	}
}

func printQueue() {
	if start {
		fmt.Println("Fila vazia!")
	} else {
		if len(queue) == 0 {
			fmt.Println("Fila vazia!")
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				fmt.Print(queue[i], "-")
			}
		}
	}
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
			if len(queue) == 0 {
				fmt.Printf("\nFila vazia\n\n")
			} else {
				retirado := queue[0]
				pop()
				fmt.Println("\nItem retirado", retirado)
				fmt.Println()
			}
		} else if i == "3" {
			fmt.Println("\nExibindo a fila")
			fmt.Println("----------------")
			printQueue()
			fmt.Printf("\n\n")
		} else {
			fmt.Printf("\nOpção inválida!\n\n")
		}
	}
}
