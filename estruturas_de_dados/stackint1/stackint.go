/*
Implementação de uma pilha estática de inteiros

1. Definição do tamanho da pilha, pois trata-se de uma pilha
   estática, ou seja, a alocação de memória para a pilha será feita
   em tempo de compilação;
2. Definição da estrutura de pilha;
3. Declaração de uma variável do tipo da pilha;
4. Função isEmpty() que verifica se a pilha está vazia;
5. Função pop() que remove um item do topo da pilha;
6. Função push() que insere um item no topo da pilha;
7. Função returnTop() que retorna o item no topo da pilha;
8. Função main() para testar as funcionalidades.

arataca89@gmail.com
20210511

Referências:
DOXSEY. Introdução a linguagem Go. 1a Edição
São Paulo: Novatec, 2016.
*/
package main

import "fmt"

const SIZE = 5

type Stack struct {
	top   int
	items [SIZE]int
}

func isEmpty(s *Stack) bool {
	if s.top == -1 {
		return true
	}
	return false
}

func pop(s *Stack) int {
	if isEmpty(s) {
		fmt.Println("Pilha vazia.")
		return 0
	}
	item := s.items[s.top]
	s.top--
	return item
}

func push(s *Stack, i int) {
	if s.top == SIZE-1 {
		fmt.Println("Pilha cheia.")
	} else {
		s.top++
		s.items[s.top] = i
	}
}

func returnTop(s *Stack) int {
	if isEmpty(s) {
		fmt.Println("Pilha vazia.")
		return 0
	}
	return s.items[s.top]
}

func exibir(s *Stack) {
	if isEmpty(s) {
		fmt.Println("Pilha vazia.")
	} else {
		fmt.Print("BASE ")
		for i := 0; i <= s.top; i++ {
			fmt.Print(s.items[i], " ")
		}
		fmt.Println("TOPO")
	}
}

func main() {
	var pilha Stack
	pilha.top = -1

	var item, opcao int

	for {
		fmt.Println("\nPilha de inteiros")
		fmt.Println("-----------------")
		fmt.Println("<< 1 >> Exibir a pilha")
		fmt.Println("<< 2 >> Inserir item na pilha.")
		fmt.Println("<< 3 >> Retirar item da pilha.")
		fmt.Println("<< 4 >> Exibir item do topo.")
		fmt.Println("<< -1 >> Sair.")
		fmt.Print("Entre com sua opcao: ")
		fmt.Scanf("%d\r", &opcao)
		fmt.Println()

		switch opcao {
		case 1:
			exibir(&pilha)
		case 2:
			fmt.Print("Entre com o valor a ser inserido: ")
			fmt.Scanf("%d\r", &item)
			push(&pilha, item)
		case 3:
			item = pop(&pilha)
			fmt.Println("Item retirado da pilha", item)
		case 4:
			fmt.Println("Item do topo da pilha", returnTop(&pilha))
		case -1:
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
